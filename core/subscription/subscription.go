package subscription

import (
	context "context"
	"fmt"
	"time"

	"github.com/odpf/siren/pkg/cortex"
)

//go:generate mockery --name=Repository -r --case underscore --with-expecter --structname SubscriptionRepository --filename subscription_repository.go --output=./mocks
type Repository interface {
	Transactor
	List(context.Context, Filter) ([]Subscription, error)
	Create(context.Context, *Subscription) error
	Get(context.Context, uint64) (*Subscription, error)
	Update(context.Context, *Subscription) error
	Delete(context.Context, uint64) error
}

type Transactor interface {
	WithTransaction(ctx context.Context) context.Context
	Rollback(ctx context.Context, err error) error
	Commit(ctx context.Context) error
}

type Receiver struct {
	ID            uint64            `json:"id"`
	Type          string            `json:"type"`
	Configuration map[string]string `json:"configuration"`
}
type Subscription struct {
	ID        uint64            `json:"id"`
	URN       string            `json:"urn"`
	Namespace uint64            `json:"namespace"`
	Receivers []Receiver        `json:"receivers"`
	Match     map[string]string `json:"match"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func (s *Subscription) ToAlertManagerReceiverConfig() []cortex.ReceiverConfig {
	if s == nil {
		return nil
	}
	amReceiverConfig := make([]cortex.ReceiverConfig, 0)
	for idx, item := range s.Receivers {
		configMapString := make(map[string]string)
		for key, value := range item.Configuration {
			strKey := fmt.Sprintf("%v", key)
			strValue := fmt.Sprintf("%v", value)

			configMapString[strKey] = strValue
		}
		newAMReceiver := cortex.ReceiverConfig{
			Receiver:      fmt.Sprintf("%s_receiverId_%d_idx_%d", s.URN, item.ID, idx),
			Match:         s.Match,
			Configuration: configMapString,
			Type:          item.Type,
		}
		amReceiverConfig = append(amReceiverConfig, newAMReceiver)
	}
	return amReceiverConfig
}
