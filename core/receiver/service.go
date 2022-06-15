package receiver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/odpf/siren/pkg/slack"
	"github.com/pkg/errors"
	goslack "github.com/slack-go/slack"
)

const (
	Slack string = "slack"
)

//go:generate mockery --name=SecureServiceProxy -r --case underscore --with-expecter --structname SecureServiceProxy --filename secure_service.go --output=./mocks
type SecureServiceProxy interface {
	ListReceivers() ([]*Receiver, error)
	CreateReceiver(*Receiver) error
	GetReceiver(uint64) (*Receiver, error)
	UpdateReceiver(*Receiver) error
	DeleteReceiver(uint64) error
	NotifyReceiver(rcv *Receiver, payloadMessage string, payloadReceiverName string, payloadReceiverType string, payloadBlock []byte) error
	Migrate() error
}

// Service handles business logic
type Service struct {
	secureService SecureServiceProxy
	slackClient   SlackClient
}

// NewService returns service struct
func NewService(secureService SecureServiceProxy, slackClient SlackClient) *Service {
	return &Service{
		secureService: secureService,
		slackClient:   slackClient,
	}
}

func (s *Service) ListReceivers() ([]*Receiver, error) {
	receivers, err := s.secureService.ListReceivers()
	if err != nil {
		return nil, err
	}
	return receivers, nil
}

func (s *Service) CreateReceiver(rcv *Receiver) error {
	if err := s.secureService.CreateReceiver(rcv); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetReceiver(id uint64) (*Receiver, error) {
	rcv, err := s.secureService.GetReceiver(id)
	if err != nil {
		return nil, err
	}

	if rcv.Type == Slack {
		token, ok := rcv.Configurations["token"].(string)
		if !ok {
			return nil, errors.New("no token found in configurations")
		}
		channels, err := s.slackClient.GetWorkspaceChannels(
			slack.CallWithContext(context.Background()),
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
	}

	return rcv, nil
}

func (s *Service) UpdateReceiver(rcv *Receiver) error {
	if err := s.secureService.UpdateReceiver(rcv); err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteReceiver(id uint64) error {
	return s.secureService.DeleteReceiver(id)
}

func (s *Service) NotifyReceiver(rcv *Receiver, payloadMessage string, payloadReceiverName string, payloadReceiverType string, payloadBlock []byte) error {
	switch rcv.Type {
	case Slack:
		blocks := goslack.Blocks{}
		if err := json.Unmarshal(payloadBlock, &blocks); err != nil {
			return fmt.Errorf("unable to parse slack block: %w", ErrInvalid)
		}

		token, ok := rcv.Configurations["token"].(string)
		if !ok {
			return fmt.Errorf("no token found in configuration: %w", ErrInvalid)
		}

		payloadMessage := &slack.Message{
			ReceiverName: payloadReceiverName,
			ReceiverType: payloadReceiverType,
			Token:        rcv.Configurations["token"].(string),
			Message:      payloadMessage,
			Blocks:       blocks,
		}
		if err := s.slackClient.Notify(payloadMessage, slack.CallWithToken(token)); err != nil {
			return fmt.Errorf("failed to notify: %w", err)
		}

	default:
		return errors.New("type not recognized")
	}
	return nil
}

func (s *Service) Migrate() error {
	return s.secureService.Migrate()
}
