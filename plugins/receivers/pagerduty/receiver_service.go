package pagerduty

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/odpf/siren/pkg/errors"
	"github.com/odpf/siren/plugins/receivers/base"
)

// ReceiverService is a receiver plugin service layer for pagerduty
type ReceiverService struct {
	base.UnimplementedReceiverService
}

// NewService returns pagerduty service struct. This service implement [receiver.Resolver] interface.
func NewReceiverService() *ReceiverService {
	return &ReceiverService{}
}

func (s *ReceiverService) PreHookTransformConfigs(ctx context.Context, configurations map[string]interface{}) (map[string]interface{}, error) {
	receiverConfig := &ReceiverConfig{}
	if err := mapstructure.Decode(configurations, receiverConfig); err != nil {
		return nil, fmt.Errorf("failed to transform configurations to receiver config: %w", err)
	}

	if err := receiverConfig.Validate(); err != nil {
		return nil, errors.ErrInvalid.WithMsgf(err.Error())
	}

	return configurations, nil
}

func (s *ReceiverService) PostHookTransformConfigs(ctx context.Context, configurations map[string]interface{}) (map[string]interface{}, error) {
	return configurations, nil
}

func (s *ReceiverService) BuildData(ctx context.Context, configurations map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func (s *ReceiverService) BuildNotificationConfig(subsConfs map[string]interface{}, receiverConfs map[string]interface{}) (map[string]interface{}, error) {
	receiverConfig := &ReceiverConfig{}
	if err := mapstructure.Decode(receiverConfs, receiverConfig); err != nil {
		return nil, fmt.Errorf("failed to transform configurations to receiver config: %w", err)
	}

	notificationConfig := NotificationConfig{
		ReceiverConfig: *receiverConfig,
	}

	return notificationConfig.AsMap(), nil
}
