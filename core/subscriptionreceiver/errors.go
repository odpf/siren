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
	ErrStr         string
}

func (err NotFoundError) Error() string {
	if err.ErrStr != "" {
		return err.ErrStr
	}
	return fmt.Sprintf("subscription with id %d and receiver with id %d not found", err.SubscriptionID, err.ReceiverID)
}
