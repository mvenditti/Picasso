package ports

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type TranscriberService interface {
	Transcribe(c context.Context, input []byte) (openai.AudioResponse, error)
}

type ImageGeneratorService interface {
	Generate(c context.Context, input string) (openai.ImageResponse, error)
}
