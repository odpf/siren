package notification

import (
	"context"
	"fmt"

	"github.com/goto/siren/core/log"
	"github.com/goto/siren/core/silence"
	"github.com/goto/siren/pkg/errors"
)

// DispatchSingleNotificationService supports subscriber routing and receiver routing at the same time
type DispatchSingleNotificationService struct {
	deps            Deps
	notifierPlugins map[string]Notifier
	routerMap       map[string]Router
}

func NewDispatchSingleNotificationService(
	deps Deps,
	notifierPlugins map[string]Notifier,
	routerMap map[string]Router,
) *DispatchSingleNotificationService {
	return &DispatchSingleNotificationService{
		deps:            deps,
		notifierPlugins: notifierPlugins,
		routerMap:       routerMap,
	}
}

func (s *DispatchSingleNotificationService) getRouter(notificationRouterKind string) (Router, error) {
	selectedRouter, exist := s.routerMap[notificationRouterKind]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported notification router kind: %q", notificationRouterKind)
	}
	return selectedRouter, nil
}

func (s *DispatchSingleNotificationService) Dispatch(ctx context.Context, ns []Notification) ([]string, error) {
	if len(ns) != 1 {
		return nil, errors.ErrInvalid.WithMsgf("direct single notification should only accept 1 notification but found %d", len(ns))
	}

	var (
		n        = ns[0]
		messages []Message
	)

	no, err := s.deps.Repository.Create(ctx, n)
	if err != nil {
		return nil, err
	}

	n.EnrichID(no.ID)

	switch n.Type {
	case TypeAlert:
		messages, err = s.dispatchByRouter(ctx, n, RouterSubscriber)
		if err != nil {
			return nil, err
		}
	case TypeEvent:
		messages, err = s.fetchMessagesEvents(ctx, n)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.ErrInternal.WithMsgf("unknown notification type %s", n.Type)
	}

	if len(messages) == 0 {
		s.deps.Logger.Info("no messages to enqueue")
		return []string{n.ID}, nil
	}

	if err := s.deps.Q.Enqueue(ctx, messages...); err != nil {
		return nil, fmt.Errorf("failed enqueuing messages: %w", err)
	}

	return []string{n.ID}, nil
}

func (s *DispatchSingleNotificationService) dispatchByRouter(ctx context.Context, n Notification, flow string) ([]Message, error) {
	if err := n.Validate(flow); err != nil {
		return nil, err
	}

	router, err := s.getRouter(flow)
	if err != nil {
		return nil, err
	}

	var (
		messages         []Message
		notificationLogs []log.Notification
		hasSilenced      bool
	)
	if s.deps.Cfg.SubscriptionV2Enabled {
		messages, notificationLogs, hasSilenced, err = router.PrepareMessageV2(ctx, n)
		if err != nil {
			return nil, err
		}
	} else {
		messages, notificationLogs, hasSilenced, err = router.PrepareMessage(ctx, n)
		if err != nil {
			return nil, err
		}
	}

	if len(messages) == 0 && len(notificationLogs) == 0 {
		return nil, fmt.Errorf("something wrong and no messages will be sent with notification: %v", n)
	}

	if err := s.deps.LogService.LogNotifications(ctx, notificationLogs...); err != nil {
		return nil, fmt.Errorf("failed logging notifications: %w", err)
	}

	// Reliability of silence feature need to be tested more
	if s.deps.Cfg.EnableSilenceFeature {
		if err := s.deps.AlertRepository.BulkUpdateSilence(ctx, n.AlertIDs, silence.Status(hasSilenced, len(messages) != 0)); err != nil {
			return nil, fmt.Errorf("failed updating silence status: %w", err)
		}
	}

	return messages, nil
}

func (s *DispatchSingleNotificationService) fetchMessagesEvents(ctx context.Context, n Notification) ([]Message, error) {
	if len(n.ReceiverSelectors) == 0 && len(n.Labels) == 0 {
		return nil, errors.ErrInvalid.WithMsgf("no receivers found")
	}

	var messages = []Message{}

	if len(n.ReceiverSelectors) != 0 {
		generatedMessages, err := s.dispatchByRouter(ctx, n, RouterReceiver)
		if err != nil {
			return nil, err
		}
		messages = append(messages, generatedMessages...)
	}

	if len(n.Labels) != 0 {
		generatedMessages, err := s.dispatchByRouter(ctx, n, RouterSubscriber)
		if err != nil {
			return nil, err
		}
		messages = append(messages, generatedMessages...)
	}
	return messages, nil
}
