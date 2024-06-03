package notification

import (
	"errors"
)

var (
	ErrNoMessage                   = errors.New("no message found")
	ErrRouteSubscriberNoMatchFound = errors.New("not matching any subscription")
)
