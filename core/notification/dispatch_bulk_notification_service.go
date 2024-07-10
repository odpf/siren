package notification

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/goto/siren/core/log"
	"github.com/goto/siren/pkg/errors"
	"github.com/goto/siren/pkg/structure"
	"github.com/mitchellh/hashstructure/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"golang.org/x/exp/maps"
)

// DispatchBulkNotificationService only supports subscriber routing and not supporting direct receiver routing
type DispatchBulkNotificationService struct {
	deps                           Deps
	notifierPlugins                map[string]Notifier
	routerMap                      map[string]Router
	metricGaugeNumBulkNotification metric.Int64Gauge
}

func NewDispatchBulkNotificationService(
	deps Deps,
	notifierPlugins map[string]Notifier,
	routerMap map[string]Router,
) *DispatchBulkNotificationService {
	metricGaugeNumBulkNotification, err := otel.Meter("github.com/goto/siren/core/notification").
		Int64Gauge("siren.notification.bulk.notification_number")
	if err != nil {
		otel.Handle(err)
	}

	return &DispatchBulkNotificationService{
		deps:                           deps,
		notifierPlugins:                notifierPlugins,
		routerMap:                      routerMap,
		metricGaugeNumBulkNotification: metricGaugeNumBulkNotification,
	}
}

func (s *DispatchBulkNotificationService) getRouter(notificationRouterKind string) (Router, error) {
	selectedRouter, exist := s.routerMap[notificationRouterKind]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported notification router kind: %q", notificationRouterKind)
	}
	return selectedRouter, nil
}

func (s *DispatchBulkNotificationService) prepareMetaMessages(ctx context.Context, ns []Notification, continueWhenError bool) (metaMessages []MetaMessage, notificationLogs []log.Notification, err error) {
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
			if errors.Is(err, ErrRouteSubscriberNoMatchFound) {
				errMessage := fmt.Sprintf("not matching any subscription for notification: %v", n)
				nJson, err := json.MarshalIndent(n, "", "  ")
				if err == nil {
					errMessage = fmt.Sprintf("not matching any subscription for notification: %s", string(nJson))
				}
				s.deps.Logger.Warn(errMessage)
				if continueWhenError {
					continue
				}
			}
			return nil, nil, err
		}

		metaMessages = append(metaMessages, generatedMetaMessages...)
		notificationLogs = append(notificationLogs, generatedNotificationLogs...)
	}

	return metaMessages, notificationLogs, nil
}

func (s *DispatchBulkNotificationService) Dispatch(ctx context.Context, ns []Notification) (notificationIDs []string, err error) {
	defer func() {
		s.instrumentNumberBulkNotification(ctx, len(ns), err)
	}()

	var (
		metaMessages     []MetaMessage
		notificationLogs []log.Notification
	)

	notifications, err := s.deps.Repository.BulkCreate(ctx, ns)
	if err != nil {
		return nil, err
	}

	for _, n := range ns {
		switch n.Type {
		case TypeAlert:
			metaMsgs, nfLogs, err := s.fetchMetaMessagesByRouter(ctx, n, RouterSubscriber)
			if err != nil {
				return nil, err
			}
			metaMessages = append(metaMessages, metaMsgs...)
			notificationLogs = append(notificationLogs, nfLogs...)
		case TypeEvent:
			metaMsgs, nfLogs, err := s.fetchMetaMessagesForEvents(ctx, n)
			if err != nil {
				return nil, err
			}
			metaMessages = append(metaMessages, metaMsgs...)
			notificationLogs = append(notificationLogs, nfLogs...)
		default:
			return nil, errors.ErrInternal.WithMsgf("unknown notification type %s", n.Type)
		}
	}

	if len(metaMessages) == 0 {
		s.deps.Logger.Info("no meta messages to process")
		return nil, nil
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

func (s *DispatchBulkNotificationService) fetchMetaMessagesByRouter(ctx context.Context, n Notification, flow string) (metaMessages []MetaMessage, notificationLogs []log.Notification, err error) {
	if err := n.Validate(flow); err != nil {
		return nil, nil, err
	}

	router, err := s.getRouter(flow)
	if err != nil {
		return nil, nil, err
	}

	// var (
	// 	messages         []Message
	// 	notificationLogs []log.Notification
	// )

	metaMessages, notificationLogs, err = router.PrepareMetaMessages(ctx, n)
	if err != nil {
		if errors.Is(err, ErrRouteSubscriberNoMatchFound) {
			errMessage := fmt.Sprintf("not matching any subscription for notification: %v", n)
			nJson, err := json.MarshalIndent(n, "", "  ")
			if err == nil {
				errMessage = fmt.Sprintf("not matching any subscription for notification: %s", string(nJson))
			}
			s.deps.Logger.Warn(errMessage)
		}
		return nil, nil, err
	}

	return metaMessages, notificationLogs, nil

	// if len(metaMessages) == 0 && len(notificationLogs) == 0 {
	// 	return nil, fmt.Errorf("something wrong and no messages will be sent with notification: %v", n)
	// }

	// if err := s.deps.LogService.LogNotifications(ctx, notificationLogs...); err != nil {
	// 	return nil, fmt.Errorf("failed logging notifications: %w", err)
	// }

	// reducedMetaMessages, err := ReduceMetaMessages(metaMessages, s.deps.Cfg.GroupBy)
	// if err != nil {
	// 	return nil, err
	// }

	// messages, err = s.RenderMessages(ctx, reducedMetaMessages)
	// if err != nil {
	// 	return nil, err
	// }
	// return messages, nil
}

func (s *DispatchBulkNotificationService) getNotifierPlugin(receiverType string) (Notifier, error) {
	notifierPlugin, exist := s.notifierPlugins[receiverType]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported receiver type: %q", receiverType)
	}
	return notifierPlugin, nil
}

func (s *DispatchBulkNotificationService) fetchMetaMessagesForEvents(ctx context.Context, n Notification) ([]MetaMessage, []log.Notification, error) {
	if len(n.ReceiverSelectors) == 0 && len(n.Labels) == 0 {
		return nil, nil, errors.ErrInvalid.WithMsgf("no receivers found")
	}

	var (
		metaMessages     = []MetaMessage{}
		notificationLogs = []log.Notification{}
	)
	if len(n.ReceiverSelectors) != 0 {
		generatedMetaMessages, nfLog, err := s.fetchMetaMessagesByRouter(ctx, n, RouterReceiver)
		if err != nil {
			return nil, nil, err
		}
		metaMessages = append(metaMessages, generatedMetaMessages...)
		notificationLogs = append(notificationLogs, nfLog...)
	}

	if len(n.Labels) != 0 {
		generatedMetaMessages, nfLog, err := s.fetchMetaMessagesByRouter(ctx, n, RouterSubscriber)
		if err != nil {
			return nil, nil, err
		}
		metaMessages = append(metaMessages, generatedMetaMessages...)
		notificationLogs = append(notificationLogs, nfLog...)
	}
	return metaMessages, notificationLogs, nil
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

func (s *DispatchBulkNotificationService) instrumentNumberBulkNotification(ctx context.Context, num int, err error) {
	s.metricGaugeNumBulkNotification.Record(ctx, int64(num), metric.WithAttributes(
		attribute.Bool("success", err == nil),
	))
}
