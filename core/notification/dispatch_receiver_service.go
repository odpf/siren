package notification

import (
	"context"

	"github.com/goto/siren/core/log"
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/pkg/errors"
)

type DispatchReceiverConfig struct {
	MaxMessagesReceiverFlow int
	MaxNumReceiverSelectors int
}

type DispatchReceiverService struct {
	conf            DispatchReceiverConfig
	receiverService ReceiverService
	templateService TemplateService
	notifierPlugins map[string]Notifier
}

func NewDispatchReceiverService(
	conf DispatchReceiverConfig,
	receiverService ReceiverService,
	templateService TemplateService,
	notifierPlugins map[string]Notifier) *DispatchReceiverService {
	return &DispatchReceiverService{
		conf:            conf,
		receiverService: receiverService,
		notifierPlugins: notifierPlugins,
		templateService: templateService,
	}
}

func (s *DispatchReceiverService) getNotifierPlugin(receiverType string) (Notifier, error) {
	notifierPlugin, exist := s.notifierPlugins[receiverType]
	if !exist {
		return nil, errors.ErrInvalid.WithMsgf("unsupported receiver type: %q", receiverType)
	}
	return notifierPlugin, nil
}

func (s *DispatchReceiverService) PrepareMessage(ctx context.Context, n Notification) ([]Message, []log.Notification, bool, error) {

	var notificationLogs []log.Notification

	if len(n.ReceiverSelectors) > s.conf.MaxNumReceiverSelectors {
		return nil, nil, false, errors.ErrInvalid.WithMsgf("number of receiver selectors should be less than or equal threshold %d", s.conf.MaxNumReceiverSelectors)
	}

	rcvs, err := s.receiverService.List(ctx, receiver.Filter{
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
			s.templateService,
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
	if messagesNum > s.conf.MaxMessagesReceiverFlow {
		return nil, nil, false, errors.ErrInvalid.WithMsgf("sending %d messages exceed max messages receiver flow threshold %d. this will spam and broadcast to %d channel. found %d receiver selectors passed, you might want to check your receiver selectors configuration", messagesNum, s.conf.MaxMessagesReceiverFlow, messagesNum, len(n.ReceiverSelectors))
	}

	return messages, notificationLogs, false, nil
}
