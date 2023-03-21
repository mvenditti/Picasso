package dependencies

import (
	"OpenAIDemo/internal/core/services/http"
	"OpenAIDemo/internal/core/services/image_generator"
	"OpenAIDemo/internal/core/services/transcriber"
	"OpenAIDemo/internal/core/services/translator"
	"OpenAIDemo/internal/handlers/artist"
)

type Definition struct {
	httpService *http.Service

	artistHandler artist.Handler
}

func NewByEnvironment() Definition {
	d := Definition{}

	d.httpService = http.NewService()

	imageGeneratorService := image_generator.NewService()
	transcriberService := transcriber.NewService()
	translatorService := translator.NewService()

	d.artistHandler = artist.NewHandler(imageGeneratorService, transcriberService, translatorService)

	return Definition{}
}
