package transcriber

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (service Service) Transcribe(input []byte) (string, error) {
	return "", nil
}
