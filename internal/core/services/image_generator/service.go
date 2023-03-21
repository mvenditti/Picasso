package image_generator

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type Service struct {
	client *openai.Client
}

func NewService(client *openai.Client) *Service {
	return &Service{client: client}
}

func (service Service) Generate(c context.Context, input string) (openai.ImageResponse, error) {
	request := openai.ImageRequest{
		Prompt:         input,
		N:              1,
		Size:           "1024x1024",
		ResponseFormat: "url",
		User:           "mvenditti",
	}

	image, err := service.client.CreateImage(c, request)

	if err != nil {

	}

	return image, nil
}
