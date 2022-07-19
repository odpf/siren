package receiver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/odpf/siren/pkg/errors"
	"github.com/odpf/siren/pkg/slack"
)

type SlackService struct {
	slackClient  SlackClient
	cryptoClient Encryptor
}

// NewService returns slack service struct
func NewSlackService(slackClient SlackClient, cryptoClient Encryptor) *SlackService {
	return &SlackService{
		slackClient:  slackClient,
		cryptoClient: cryptoClient,
	}
}

func (s *SlackService) Notify(ctx context.Context, rcv *Receiver, payloadMessage NotificationMessage) error {
	token, ok := rcv.Configurations["token"].(string)
	if !ok {
		return errors.ErrInvalid.WithMsgf("no token in configurations found")
	}

	sm, err := payloadMessage.ToSlackMessage()
	if err != nil {
		return err
	}

	if err := s.slackClient.Notify(ctx, sm, slack.CallWithToken(token)); err != nil {
		return fmt.Errorf("error calling slack notify: %w", err)
	}

	return nil
}

func (s *SlackService) Encrypt(r *Receiver) error {
	token, ok := r.Configurations["token"].(string)
	if !ok {
		return errors.ErrInvalid.WithMsgf("no token in configurations found")
	}
	chiperText, err := s.cryptoClient.Encrypt(token)
	if err != nil {
		return err
	}
	r.Configurations["token"] = chiperText

	return nil
}

func (s *SlackService) Decrypt(r *Receiver) error {
	cipherText, ok := r.Configurations["token"].(string)
	if !ok {
		return errors.ErrInvalid.WithMsgf("no token in configurations found")
	}
	token, err := s.cryptoClient.Decrypt(cipherText)
	if err != nil {
		return err
	}
	r.Configurations["token"] = token
	return nil
}

func (s *SlackService) PopulateReceiver(ctx context.Context, rcv *Receiver) (*Receiver, error) {
	token, ok := rcv.Configurations["token"].(string)
	if !ok {
		return nil, errors.ErrInvalid.WithMsgf("no token in configurations found")
	}

	channels, err := s.slackClient.GetWorkspaceChannels(
		ctx,
		slack.CallWithToken(token),
	)
	if err != nil {
		return nil, fmt.Errorf("could not get channels: %w", err)
	}

	data, err := json.Marshal(channels)
	if err != nil {
		// this is very unlikely to return error since we have an explicitly defined type of channels
		return nil, fmt.Errorf("invalid channels: %w", err)
	}

	rcv.Data = make(map[string]interface{})
	rcv.Data["channels"] = string(data)

	return rcv, nil
}

func (s *SlackService) ValidateConfiguration(rcv *Receiver) error {
	if rcv == nil {
		return errors.New("receiver to validate is nil")
	}

	_, err := rcv.Configurations.GetString("client_id")
	if err != nil {
		return err
	}

	_, err = rcv.Configurations.GetString("client_secret")
	if err != nil {
		return err
	}

	_, err = rcv.Configurations.GetString("auth_code")
	if err != nil {
		return err
	}

	return nil
}

func (s *SlackService) GetSubscriptionConfig(subsConfs map[string]string, receiverConfs Configurations) (map[string]string, error) {
	mapConf := make(map[string]string)
	if _, ok := subsConfs["channel_name"]; !ok {
		return nil, errors.New("subscription receiver config 'channel_name' was missing")
	}
	mapConf["channel_name"] = subsConfs["channel_name"]
	if val, ok := receiverConfs["token"]; ok {
		if mapConf["token"], ok = val.(string); !ok {
			return nil, errors.New("token config from receiver should be in string")
		}
	}
	return mapConf, nil
}
