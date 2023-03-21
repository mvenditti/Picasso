package translator

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (service Service) Translate(input string) (string, error) {
	return "", nil
}
