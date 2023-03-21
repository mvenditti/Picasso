package dependencies

import (
	"Picasso/internal/core/services/image_generator"
	"Picasso/internal/core/services/transcriber"
	"Picasso/internal/handlers/artist"
	"Picasso/pkg/constants"
	"github.com/sashabaranov/go-openai"
	"os"
)

type Definition struct {
	ArtistHandler *artist.Handler
}

func NewByEnvironment() Definition {
	d := Definition{}

	apiKey := os.Getenv(constants.ApiKey)
	if apiKey == "" {
		panic("Failed to start up server due to empty API Key, please be aware to set API_KEY environment variable")
	}

	client := openai.NewClient(apiKey)

	imageGeneratorService := image_generator.NewService(client)
	transcriberService := transcriber.NewService(client)

	d.ArtistHandler = artist.NewHandler(imageGeneratorService, transcriberService)

	return d
}
