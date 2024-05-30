package notification

import (
	"time"
)

type MetaMessage struct {
	ReceiverID       uint64
	SubscriptionIDs  []uint64
	ReceiverType     string
	NotificationIDs  []string
	NotificationType string
	ReceiverConfigs  map[string]any
	Data             map[string]any
	ValidDuration    time.Duration
	Template         string
	Labels           map[string]string
	MergedLabels     map[string][]string
}
