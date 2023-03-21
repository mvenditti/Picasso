package http

import (
	"errors"
	"fmt"
	"github.com/mercadolibre/fury_go-core/pkg/rusty"
)

type Service map[string]Endpoint

type Endpoint struct {
	rustyEndpoint rusty.Endpoint
}

func NewService() *Service {
	return &Service{}
}

func (s Service) GetEndpoint(key string) (*Endpoint, error) {
	endpoint, found := s[key]

	if !found {
		err := errors.New(fmt.Sprintf("Couldn't fetch endpoint with key name %s", key))
		return nil, err
	}

	return &endpoint, nil
}

func (s Service) SetEndpoint(key string, endpoint Endpoint) {
	s[key] = endpoint
}
