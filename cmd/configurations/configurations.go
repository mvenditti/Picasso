package configurations

import (
	"Picasso/pkg/constants"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

func InitOpenAIClient(environment string) *openai.Client {
	baseURL := constants.ProductiveProxyURL
	if environment == constants.Local {
		baseURL = constants.TestingProxyURL
	}

	return openai.NewClientWithConfig(openai.ClientConfig{
		HTTPClient: &http.Client{},
		BaseURL:    baseURL,
	})
}
