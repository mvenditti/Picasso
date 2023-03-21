package transcriber

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

func (service Service) Transcribe(c context.Context, input []byte) (openai.AudioResponse, error) {
	request := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: "/Users/mvenditti/go/src/Picasso/files/audio.mp3",
		Language: "es",
	}

	response, err := service.client.CreateTranscription(c, request)

	if err != nil {

	}

	return response, nil
}
