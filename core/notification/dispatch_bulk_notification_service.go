package notification

import (
	"context"
	"fmt"

	"github.com/goto/siren/core/log"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/structure"
	"github.com/mitchellh/hashstructure/v2"
	"golang.org/x/exp/maps"
)

// DispatchBulkNotificationService only supports subscriber routing and not supporting direct receiver routing
type DispatchBulkNotificationService struct {
	deps            Deps
	notifierPlugins map[string]Notifier
	routerMap       map[string]Router
}

func NewDispatchBulkNotificationService(
	deps Deps,
	notifierPlugins map[string]Notifier,
	routerMap map[string]Router,
) *DispatchBulkNotificationService {
	return &DispatchBulkNotificationService{
		deps:            deps,
		notifierPlugins: notifierPlugins,
		routerMap:       routerMap,
	}
}

func (s *DispatchBulkNotificationService) getRouter(notificationRouterKind string) (Router, error) {
	selectedRouter, exist := s.routerMap[notificationRouterKind]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported notification router kind: %q", notificationRouterKind)
	}
	return selectedRouter, nil
}

func (s *DispatchBulkNotificationService) prepareMetaMessages(ctx context.Context, ns []Notification) ([]MetaMessage, []log.Notification, error) {
	var (
		metaMessages     []MetaMessage
		notificationLogs []log.Notification
	)
	for _, n := range ns {
		if err := n.Validate(RouterSubscriber); err != nil {
			return nil, nil, err
		}

		router, err := s.getRouter(RouterSubscriber)
		if err != nil {
			return nil, nil, err
		}

		generatedMetaMessages, generatedNotificationLogs, err := router.PrepareMetaMessages(ctx, n)
		if err != nil {
			return nil, nil, err
		}

		metaMessages = append(metaMessages, generatedMetaMessages...)
		notificationLogs = append(notificationLogs, generatedNotificationLogs...)
	}

	return metaMessages, notificationLogs, nil
}

func (s *DispatchBulkNotificationService) Dispatch(ctx context.Context, ns []Notification) ([]string, error) {
	var (
		notificationIDs  []string
		metaMessages     []MetaMessage
		notificationLogs []log.Notification
	)

	notifications, err := s.deps.Repository.BulkCreate(ctx, ns)
	if err != nil {
		return nil, err
	}

	metaMessages, notificationLogs, err = s.prepareMetaMessages(ctx, notifications)
	if err != nil {
		return nil, err
	}

	if err := s.deps.LogService.LogNotifications(ctx, notificationLogs...); err != nil {
		return nil, fmt.Errorf("failed logging notifications: %w", err)
	}

	reducedMetaMessages, err := ReduceMetaMessages(metaMessages, s.deps.Cfg.GroupBy)
	if err != nil {
		return nil, err
	}

	messages, err := s.RenderMessages(ctx, reducedMetaMessages)
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		s.deps.Logger.Info("no messages to enqueue")
		return nil, nil
	}

	if err := s.deps.Q.Enqueue(ctx, messages...); err != nil {
		return nil, fmt.Errorf("failed enqueuing messages: %w", err)
	}

	for _, n := range notifications {
		notificationIDs = append(notificationIDs, n.ID)
	}

	return notificationIDs, nil
}

func (s *DispatchBulkNotificationService) getNotifierPlugin(receiverType string) (Notifier, error) {
	notifierPlugin, exist := s.notifierPlugins[receiverType]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported receiver type: %q", receiverType)
	}
	return notifierPlugin, nil
}

func (s *DispatchBulkNotificationService) RenderMessages(ctx context.Context, metaMessages []MetaMessage) (messages []Message, err error) {
	for _, mm := range metaMessages {
		notifierPlugin, err := s.getNotifierPlugin(mm.ReceiverType)
		if err != nil {
			return nil, err
		}

		message, err := InitMessageByMetaMessage(
			ctx,
			notifierPlugin,
			s.deps.TemplateService,
			mm,
			InitWithExpiryDuration(mm.ValidDuration),
		)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}
	return messages, nil
}

func ReduceMetaMessages(metaMessages []MetaMessage, groupBy []string) ([]MetaMessage, error) {
	var (
		hashedMetaMessagesMap = map[uint64]MetaMessage{}
	)
	for _, mm := range metaMessages {

		groupedLabels := structure.BuildGroupLabels(mm.Labels, groupBy)
		groupedLabels["_receiver.ID"] = fmt.Sprintf("%d", mm.ReceiverID)
		groupedLabels["_notification.template"] = mm.Template

		hash, err := hashstructure.Hash(groupedLabels, hashstructure.FormatV2, nil)
		if err != nil {
			return nil, fmt.Errorf("cannot get hash from metamessage %v", mm)
		}

		if _, ok := hashedMetaMessagesMap[hash]; !ok {
			if mm.MergedLabels == nil {
				mm.MergedLabels = map[string][]string{}
				for k, v := range mm.Labels {
					mm.MergedLabels[k] = append(mm.MergedLabels[k], v)
				}
			}
			hashedMetaMessagesMap[hash] = mm

		} else {
			hashedMetaMessagesMap[hash] = MergeMetaMessage(mm, hashedMetaMessagesMap[hash])
		}

	}
	return maps.Values(hashedMetaMessagesMap), nil
}

func MergeMetaMessage(from MetaMessage, to MetaMessage) MetaMessage {
	var output = to
	for k, v := range from.Labels {
		output.MergedLabels[k] = append(output.MergedLabels[k], v)
	}
	output.NotificationIDs = append(output.NotificationIDs, from.NotificationIDs...)
	output.SubscriptionIDs = append(output.SubscriptionIDs, from.SubscriptionIDs...)
	return output
}
