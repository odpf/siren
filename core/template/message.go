package template

import (
	"github.com/goto/siren/core/receiver"
	"github.com/goto/siren/pkg/errors"
)

func MessageContentByReceiverType(messagesTemplate []Message, receiverType string) (string, error) {
	var messageTemplateMap = make(map[string]string)

	for _, msgTemplate := range messagesTemplate {
		messageTemplateMap[msgTemplate.ReceiverType] = msgTemplate.Content
	}

	// slack and slack_channel could be use interchangeably
	receiverTypeKey := receiverType
	if receiverType == receiver.TypeSlackChannel {
		receiverTypeKey = receiver.TypeSlack
	}

	content, ok := messageTemplateMap[receiverTypeKey]
	if !ok {
		errors.ErrInvalid.WithCausef("can't found template of receiver type %s", receiverType)
	}

	if content == "" {
		return "", errors.ErrInvalid.WithCausef("%s template is empty", receiverType)
	}

	return content, nil
}
