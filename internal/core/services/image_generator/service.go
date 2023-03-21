package image_generator

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (service Service) Generate(input string) ([]byte, error) {
	return []byte{}, nil
}
