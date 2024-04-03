package notification

type Filter struct {
	Type             string
	Template         string
	Labels           map[string]string
	ReceiverSelector map[string]string
}
