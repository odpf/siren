package subscriptionreceiver

import (
	"time"
)

type Relation struct {
	ID             uint64            `json:"id"`
	SubscriptionID uint64            `json:"subscription_id"`
	ReceiverID     uint64            `json:"receiver_id"`
	Labels         map[string]string `json:"labels"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
	DeletedAt      time.Time         `json:"deleted_at"`
}
