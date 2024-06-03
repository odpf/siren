package subscriptionreceiver

type Filter struct {
	SubscriptionIDs []uint64
	ReceiverID      uint64
	Labels          map[string]string
	Deleted         bool
}

type DeleteFilter struct {
	Pair           []Relation
	SubscriptionID uint64
}
