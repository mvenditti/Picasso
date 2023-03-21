package artist

import (
	"OpenAIDemo/internal/core/domain/artist"
	"OpenAIDemo/internal/core/ports"
	"encoding/json"
	"github.com/mercadolibre/fury_go-core/pkg/web"
	"net/http"
)

type Handler struct {
	imageGeneratorService ports.ImageGeneratorService
	transcriberService    ports.TranscriberService
	translatorService     ports.TranslatorService
}

func NewHandler(imageGeneratorService ports.ImageGeneratorService, transcriberService ports.TranscriberService, translatorService ports.TranslatorService) Handler {
	return Handler{
		imageGeneratorService: imageGeneratorService,
		transcriberService:    transcriberService,
		translatorService:     translatorService,
	}
}

func (handler Handler) Paint(r *http.Request, w http.ResponseWriter) error {
	var command artist.Command
	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {

	}

	transcription, err := handler.transcriberService.Transcribe(command.Input)
	if err != nil {

	}

	translated, err := handler.translatorService.Translate(transcription)
	if err != nil {

	}

	image, err := handler.imageGeneratorService.Generate(translated)
	if err != nil {

	}

	response := artist.Painting{Data: image}
	return web.EncodeJSON(w, response, http.StatusCreated)
}
