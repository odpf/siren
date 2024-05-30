package notification

import (
	"fmt"

	sirenv1beta1 "github.com/goto/siren/proto/gotocompany/siren/v1beta1"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Message) ToV1beta1Proto() (*sirenv1beta1.NotificationMessage, error) {
	details, err := structpb.NewStruct(m.Details)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch notification message details: %w", err)
	}
	configs, err := structpb.NewStruct(m.Configs)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch notification message details: %w", err)
	}

	return &sirenv1beta1.NotificationMessage{
		Id:              m.ID,
		NotificationIds: m.NotificationIDs,
		Status:          m.Status.String(),
		ReceiverType:    m.ReceiverType,
		Configs:         configs,
		Details:         details,
		LastError:       m.LastError,
		MaxTries:        uint64(m.MaxTries),
		TryCount:        uint64(m.TryCount),
		Retryable:       m.Retryable,
		ExpiredAt:       timestamppb.New(m.ExpiredAt),
		CreatedAt:       timestamppb.New(m.CreatedAt),
		UpdatedAt:       timestamppb.New(m.UpdatedAt),
	}, nil
}
