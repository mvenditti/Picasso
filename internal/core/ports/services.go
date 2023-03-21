package ports

import "OpenAIDemo/internal/core/services/http"

type HttpService interface {
	GetEndpoint(key string) (*http.Endpoint, error)
	SetEndpoint(key string, endpoint http.Endpoint)
}

type TranslatorService interface {
	Translate(input string) (string, error)
}

type TranscriberService interface {
	Transcribe(input []byte) (string, error)
}

type ImageGeneratorService interface {
	Generate(input string) ([]byte, error)
}
