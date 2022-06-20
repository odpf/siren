package receiver

type HTTPService struct{}

// NewHTTPService returns slack service struct
func NewHTTPService() *HTTPService {
	return &HTTPService{}
}

func (s *HTTPService) Notify(rcv *Receiver, payloadMessage NotificationMessage) error {
	return nil
}

func (s *HTTPService) Encrypt(r *Receiver) error {
	return nil
}

func (s *HTTPService) Decrypt(r *Receiver) error {
	return nil
}

func (s *HTTPService) PopulateReceiver(rcv *Receiver) (*Receiver, error) {
	return rcv, nil
}

func (s *HTTPService) ValidateConfiguration(configurations Configurations) error {
	_, err := configurations.GetString("url")
	if err != nil {
		return err
	}
	return nil
}
