package notification

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/goto/siren/core/log"
	"github.com/goto/siren/core/silence"
	"github.com/goto/siren/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const (
	metricRouterSubscriberAttributePrepareMessage = "preparemessage.status"
	metricRouterSubscriberAttributePrepareSuccess = "preparemessage.success"

	metricRouterSubscriberStatusMatchError       = "MATCH_ERROR"
	metricRouterSubscriberStatusMatchNotFound    = "MATCH_NOT_FOUND"
	metricRouterSubscriberStatusMessageInitError = "MESSAGE_INIT_ERROR"
	metricRouterSubscriberStatusNotifierError    = "NOTIFIER_ERROR"
	metricRouterSubscriberStatusSuccess          = "SUCCESS"
)

type RouterSubscriberService struct {
	deps                          Deps
	notifierPlugins               map[string]Notifier
	metricCounterRouterSubscriber metric.Int64Counter
}

func NewRouterSubscriberService(
	deps Deps,
	notifierPlugins map[string]Notifier,
) *RouterSubscriberService {
	metricCounterRouterSubscriber, err := otel.Meter("github.com/goto/siren/core/notification").
		Int64Counter("siren.notification.dispatch.subscriber")
	if err != nil {
		otel.Handle(err)
	}
	return &RouterSubscriberService{
		deps:                          deps,
		notifierPlugins:               notifierPlugins,
		metricCounterRouterSubscriber: metricCounterRouterSubscriber,
	}
}

func (s *RouterSubscriberService) getNotifierPlugin(receiverType string) (Notifier, error) {
	notifierPlugin, exist := s.notifierPlugins[receiverType]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported receiver type: %q", receiverType)
	}
	return notifierPlugin, nil
}

func (s *RouterSubscriberService) PrepareMessage(ctx context.Context, n Notification) ([]Message, []log.Notification, bool, error) {

	var (
		messages         = make([]Message, 0)
		notificationLogs []log.Notification
		hasSilenced      bool
	)

	subscriptions, err := s.deps.SubscriptionService.MatchByLabels(ctx, n.NamespaceID, n.Labels)
	if err != nil {
		return nil, nil, false, err
	}

	if len(subscriptions) == 0 {
		return nil, nil, false, errors.ErrInvalid.WithMsgf("not matching any subscription")
	}

	for _, sub := range subscriptions {

		if len(sub.Receivers) == 0 {
			s.deps.Logger.Warn(fmt.Sprintf("invalid subscription with id %d, no receiver found", sub.ID))
			continue
		}

		var silences []silence.Silence

		// Reliability of silence feature need to be tested more
		if s.deps.Cfg.EnableSilenceFeature {
			// try silencing by labels
			silences, err = s.deps.SilenceService.List(ctx, silence.Filter{
				NamespaceID:       n.NamespaceID,
				SubscriptionMatch: sub.Match,
			})
			if err != nil {
				return nil, nil, false, err
			}
		}

		if len(silences) != 0 {
			hasSilenced = true

			var silenceIDs []string
			for _, sil := range silences {
				silenceIDs = append(silenceIDs, sil.ID)
			}

			notificationLogs = append(notificationLogs, log.Notification{
				NamespaceID:    n.NamespaceID,
				NotificationID: n.ID,
				SubscriptionID: sub.ID,
				AlertIDs:       n.AlertIDs,
				SilenceIDs:     silenceIDs,
			})

			s.deps.Logger.Info(fmt.Sprintf("notification '%s' of alert ids '%v' is being silenced by labels '%v'", n.ID, n.AlertIDs, silences))
			continue
		}

		// Reliability of silence feature need to be tested more
		if s.deps.Cfg.EnableSilenceFeature {
			// subscription not being silenced by label
			silences, err = s.deps.SilenceService.List(ctx, silence.Filter{
				NamespaceID:    n.NamespaceID,
				SubscriptionID: sub.ID,
			})
			if err != nil {
				return nil, nil, false, err
			}
		}

		silencedReceiversMap, validReceivers, err := sub.SilenceReceivers(silences)
		if err != nil {
			return nil, nil, false, errors.ErrInvalid.WithMsgf(err.Error())
		}

		if len(silencedReceiversMap) != 0 {
			hasSilenced = true

			for rcvID, sils := range silencedReceiversMap {
				var silenceIDs []string
				for _, sil := range sils {
					silenceIDs = append(silenceIDs, sil.ID)
				}

				notificationLogs = append(notificationLogs, log.Notification{
					NamespaceID:    n.NamespaceID,
					NotificationID: n.ID,
					SubscriptionID: sub.ID,
					ReceiverID:     rcvID,
					AlertIDs:       n.AlertIDs,
					SilenceIDs:     silenceIDs,
				})
			}
		}

		for _, rcv := range validReceivers {
			notifierPlugin, err := s.getNotifierPlugin(rcv.Type)
			if err != nil {
				return nil, nil, false, err
			}

			message, err := InitMessage(
				ctx,
				notifierPlugin,
				s.deps.TemplateService,
				n,
				rcv.Type,
				rcv.Configuration,
				InitWithExpiryDuration(n.ValidDuration),
			)
			if err != nil {
				return nil, nil, false, err
			}

			messages = append(messages, message)
			notificationLogs = append(notificationLogs, log.Notification{
				NamespaceID:    n.NamespaceID,
				NotificationID: n.ID,
				SubscriptionID: sub.ID,
				ReceiverID:     rcv.ID,
				AlertIDs:       n.AlertIDs,
			})
		}
	}

	return messages, notificationLogs, hasSilenced, nil
}

func (s *RouterSubscriberService) PrepareMessageV2(ctx context.Context, n Notification) (messages []Message, notificationLogs []log.Notification, hasSilenced bool, err error) {
	var metricStatus = metricRouterSubscriberStatusSuccess

	messages = make([]Message, 0)

	defer func() {
		s.instrumentDispatchSubscription(ctx, metricRouterSubscriberAttributePrepareMessage, metricStatus, err)
	}()

	receiversView, err := s.deps.SubscriptionService.MatchByLabelsV2(ctx, n.NamespaceID, n.Labels)
	if err != nil {
		metricStatus = metricRouterSubscriberStatusMatchError
		return nil, nil, false, err
	}

	if len(receiversView) == 0 {
		metricStatus = metricRouterSubscriberStatusMatchNotFound
		return nil, nil, false, errors.ErrInvalid.WithMsgf("not matching any subscription")
	}

	for _, rcv := range receiversView {
		notifierPlugin, err := s.getNotifierPlugin(rcv.Type)
		if err != nil {
			metricStatus = metricRouterSubscriberStatusNotifierError
			return nil, nil, false, err
		}

		message, err := InitMessage(
			ctx,
			notifierPlugin,
			s.deps.TemplateService,
			n,
			rcv.Type,
			rcv.Configurations,
			InitWithExpiryDuration(n.ValidDuration),
		)
		if err != nil {
			metricStatus = metricRouterSubscriberStatusMessageInitError
			return nil, nil, false, err
		}

		messages = append(messages, message)
		notificationLogs = append(notificationLogs, log.Notification{
			NamespaceID:    n.NamespaceID,
			NotificationID: n.ID,
			SubscriptionID: rcv.SubscriptionID,
			ReceiverID:     rcv.ID,
			AlertIDs:       n.AlertIDs,
		})
	}

	return messages, notificationLogs, hasSilenced, nil
}

func (s *RouterSubscriberService) instrumentDispatchSubscription(ctx context.Context, attributeKey, attributeValue string, err error) {
	s.metricCounterRouterSubscriber.Add(ctx, 1, metric.WithAttributes(
		attribute.String(attributeKey, attributeValue),
		attribute.Bool("operation.success", err == nil),
	))
}

func (s *RouterSubscriberService) PrepareMetaMessages(ctx context.Context, n Notification) (metaMessages []MetaMessage, notificationLogs []log.Notification, err error) {
	var metricStatus = metricRouterSubscriberStatusSuccess

	defer func() {
		s.instrumentDispatchSubscription(ctx, metricRouterSubscriberAttributePrepareMessage, metricStatus, err)
	}()

	receiversView, err := s.deps.SubscriptionService.MatchByLabelsV2(ctx, n.NamespaceID, n.Labels)
	if err != nil {
		metricStatus = metricRouterSubscriberStatusMatchError
		return nil, nil, err
	}

	if len(receiversView) == 0 {
		metricStatus = metricRouterSubscriberStatusMatchNotFound
		errMessage := fmt.Sprintf("not matching any subscription for notification: %v", n)
		nJson, err := json.MarshalIndent(n, "", "  ")
		if err == nil {
			errMessage = fmt.Sprintf("not matching any subscription for notification: %s", string(nJson))
		}
		return nil, nil, errors.ErrInvalid.WithMsgf(errMessage)
	}

	for _, rcv := range receiversView {
		metaMessages = append(metaMessages, n.MetaMessage(rcv))

		// messages = append(messages, message)
		notificationLogs = append(notificationLogs, log.Notification{
			NamespaceID:    n.NamespaceID,
			NotificationID: n.ID,
			SubscriptionID: rcv.SubscriptionID,
			ReceiverID:     rcv.ID,
			AlertIDs:       n.AlertIDs,
		})
	}

	return metaMessages, notificationLogs, nil
}
