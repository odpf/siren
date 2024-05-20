package subscriptionreceiver

import (
	"errors"
	"fmt"
)

var (
	ErrDuplicate = errors.New("subscription id and receiver id already registered")
	ErrRelation  = errors.New("subscription or receiver id does not exist")
)

type NotFoundError struct {
	SubscriptionID uint64
	ReceiverID     uint64
}

func (err NotFoundError) Error() string {
	if err.SubscriptionID != 0 {
		return fmt.Sprintf("subscription with id %d not found", err.SubscriptionID)
	}
	if err.ReceiverID != 0 {
		return fmt.Sprintf("receiver with id %d not found", err.ReceiverID)
	}

	return "subscription receiver pair not found"
}
