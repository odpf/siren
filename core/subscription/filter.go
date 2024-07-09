package subscription

type Filter struct {
	NamespaceID                uint64
	Match                      map[string]string
	NotificationMatch          map[string]string
	Metadata                   map[string]any
	SilenceID                  string
	IDs                        []int64
	SubscriptionReceiverLabels map[string]string
	ReceiverID                 uint64
	WithSubscriptionReceiver   bool
}
