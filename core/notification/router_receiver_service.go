package notification

import (
	"context"

	"github.com/goto/siren/core/log"
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/core/subscription"
	"github.com/goto/siren/pkg/errors"
)

type RouterReceiverService struct {
	deps            Deps
	notifierPlugins map[string]Notifier
}

func NewRouterReceiverService(
	deps Deps,
	notifierPlugins map[string]Notifier,
) *RouterReceiverService {
	return &RouterReceiverService{
		deps:            deps,
		notifierPlugins: notifierPlugins,
	}
}

func (s *RouterReceiverService) getNotifierPlugin(receiverType string) (Notifier, error) {
	notifierPlugin, exist := s.notifierPlugins[receiverType]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported receiver type: %q", receiverType)
	}
	return notifierPlugin, nil
}

func (s *RouterReceiverService) PrepareMetaMessages(ctx context.Context, n Notification) (metaMessages []MetaMessage, notificationLogs []log.Notification, err error) {
	if len(n.ReceiverSelectors) > s.deps.Cfg.MaxNumReceiverSelectors {
		return nil, nil, errors.ErrInvalid.WithMsgf("number of receiver selectors should be less than or equal threshold %d", s.deps.Cfg.MaxNumReceiverSelectors)
	}

	rcvs, err := s.deps.ReceiverService.List(ctx, receiver.Filter{
		MultipleLabels: n.ReceiverSelectors,
		Expanded:       true,
	})
	if err != nil {
		return nil, nil, err
	}

	if len(rcvs) == 0 {
		return nil, nil, errors.ErrNotFound
	}

	for _, rcv := range rcvs {
		var rcvView = &subscription.ReceiverView{}
		rcvView.FromReceiver(rcv)
		metaMessages = append(metaMessages, n.MetaMessage(*rcvView))

		notificationLogs = append(notificationLogs, log.Notification{
			NamespaceID:    n.NamespaceID,
			NotificationID: n.ID,
			ReceiverID:     rcv.ID,
			AlertIDs:       n.AlertIDs,
		})
	}

	var metaMessagesNum = len(metaMessages)
	if metaMessagesNum > s.deps.Cfg.MaxMessagesReceiverFlow {
		return nil, nil, errors.ErrInvalid.WithMsgf("sending %d messages exceed max messages receiver flow threshold %d. this will spam and broadcast to %d channel. found %d receiver selectors passed, you might want to check your receiver selectors configuration", metaMessagesNum, s.deps.Cfg.MaxMessagesReceiverFlow, metaMessagesNum, len(n.ReceiverSelectors))
	}

	return metaMessages, notificationLogs, nil
}

func (s *RouterReceiverService) PrepareMessageV2(ctx context.Context, n Notification) ([]Message, []log.Notification, bool, error) {
	return s.PrepareMessage(ctx, n)
}

func (s *RouterReceiverService) PrepareMessage(ctx context.Context, n Notification) ([]Message, []log.Notification, bool, error) {

	var notificationLogs []log.Notification

	if len(n.ReceiverSelectors) > s.deps.Cfg.MaxNumReceiverSelectors {
		return nil, nil, false, errors.ErrInvalid.WithMsgf("number of receiver selectors should be less than or equal threshold %d", s.deps.Cfg.MaxNumReceiverSelectors)
	}

	rcvs, err := s.deps.ReceiverService.List(ctx, receiver.Filter{
		MultipleLabels: n.ReceiverSelectors,
		Expanded:       true,
	})
	if err != nil {
		return nil, nil, false, err
	}

	if len(rcvs) == 0 {
		return nil, nil, false, errors.ErrNotFound
	}

	var messages []Message

	for _, rcv := range rcvs {
		notifierPlugin, err := s.getNotifierPlugin(rcv.Type)
		if err != nil {
			return nil, nil, false, errors.ErrInvalid.WithMsgf("invalid receiver type: %s", err.Error())
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
			return nil, nil, false, err
		}

		messages = append(messages, message)
		notificationLogs = append(notificationLogs, log.Notification{
			NamespaceID:    n.NamespaceID,
			NotificationID: n.ID,
			ReceiverID:     rcv.ID,
			AlertIDs:       n.AlertIDs,
		})
	}

	var messagesNum = len(messages)
	if messagesNum > s.deps.Cfg.MaxMessagesReceiverFlow {
		return nil, nil, false, errors.ErrInvalid.WithMsgf("sending %d messages exceed max messages receiver flow threshold %d. this will spam and broadcast to %d channel. found %d receiver selectors passed, you might want to check your receiver selectors configuration", messagesNum, s.deps.Cfg.MaxMessagesReceiverFlow, messagesNum, len(n.ReceiverSelectors))
	}

	return messages, notificationLogs, false, nil
}
