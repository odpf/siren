package notification

import (
	"context"
	"time"

	saltlog "github.com/goto/salt/log"

	"github.com/goto/siren/core/log"
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/core/silence"
	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/core/template"
	"github.com/goto/siren/pkg/errors"
)

type Router interface {
	PrepareMessage(ctx context.Context, n Notification) ([]Message, []log.Notification, bool, error)
	PrepareMessageV2(ctx context.Context, n Notification) ([]Message, []log.Notification, bool, error)
	PrepareMetaMessages(ctx context.Context, n Notification) (metaMessages []MetaMessage, notificationLogs []log.Notification, err error)
}

type Dispatcher interface {
	Dispatch(ctx context.Context, ns []Notification) ([]string, error)
}

type SubscriptionService interface {
	MatchByLabels(ctx context.Context, namespaceID uint64, labels map[string]string) ([]subscription.Subscription, error)
	MatchByLabelsV2(ctx context.Context, namespaceID uint64, labels map[string]string) ([]subscription.ReceiverView, error)
}

type ReceiverService interface {
	List(ctx context.Context, flt receiver.Filter) ([]receiver.Receiver, error)
}

type SilenceService interface {
	List(ctx context.Context, filter silence.Filter) ([]silence.Silence, error)
}

type AlertRepository interface {
	BulkUpdateSilence(context.Context, []int64, string) error
}

type LogService interface {
	LogNotifications(ctx context.Context, nlogs ...log.Notification) error
}

type TemplateService interface {
	GetByName(ctx context.Context, name string) (*template.Template, error)
}

// Service is a service for notification domain
type Service struct {
	deps               Deps
	dispatchServiceMap map[string]Dispatcher
}

type Deps struct {
	Cfg                   Config
	Logger                saltlog.Logger
	Repository            Repository
	Q                     Queuer
	IdempotencyRepository IdempotencyRepository
	AlertRepository       AlertRepository
	LogService            LogService
	ReceiverService       ReceiverService
	TemplateService       TemplateService
	SubscriptionService   SubscriptionService
	SilenceService        SilenceService
}

// NewService creates a new notification service
func NewService(
	deps Deps,
	dispatchServiceMap map[string]Dispatcher,
) *Service {
	return &Service{
		deps:               deps,
		dispatchServiceMap: dispatchServiceMap,
	}
}

func (s *Service) Dispatch(ctx context.Context, ns []Notification, dispatcherKind string) ([]string, error) {
	ctx = s.deps.Repository.WithTransaction(ctx)
	selectedDispatcher, exist := s.dispatchServiceMap[dispatcherKind]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported notification dispatcher: %q", dispatcherKind)
	}
	ids, err := selectedDispatcher.Dispatch(ctx, ns)
	if err != nil {
		if err := s.deps.Repository.Rollback(ctx, err); err != nil {
			return nil, err
		}
		return nil, err
	}

	if err := s.deps.Repository.Commit(ctx); err != nil {
		return nil, err
	}

	return ids, nil
}

func (s *Service) CheckIdempotency(ctx context.Context, scope, key string) (string, error) {
	idempt, err := s.deps.IdempotencyRepository.Check(ctx, scope, key)
	if err != nil {
		return "", err
	}

	return idempt.NotificationID, nil
}

func (s *Service) InsertIdempotency(ctx context.Context, scope, key, notificationID string) error {
	if _, err := s.deps.IdempotencyRepository.Create(ctx, scope, key, notificationID); err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveIdempotencies(ctx context.Context, TTL time.Duration) error {
	return s.deps.IdempotencyRepository.Delete(ctx, IdempotencyFilter{
		TTL: TTL,
	})
}

func (s *Service) ListNotificationMessages(ctx context.Context, notificationID string) ([]Message, error) {
	messages, err := s.deps.Q.ListMessages(ctx, notificationID)
	if err != nil {
		return nil, err
	}

	messages = s.discardSecrets(messages)

	return messages, nil
}

// TODO might want to do smarter way to discard secrets
func (s *Service) discardSecrets(messages []Message) []Message {
	newMessages := make([]Message, 0)

	for _, msg := range messages {
		newMsg := msg
		cfg := newMsg.Configs
		// slack token
		delete(cfg, "token")
		// pagerduty service key
		delete(cfg, "service_key")
		newMsg.Configs = cfg
		newMessages = append(newMessages, newMsg)
	}

	return newMessages
}

func (s *Service) List(ctx context.Context, flt Filter) ([]Notification, error) {
	notifications, err := s.deps.Repository.List(ctx, flt)
	if err != nil {
		return nil, err
	}

	return notifications, err
}
