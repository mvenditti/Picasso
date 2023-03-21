package artist

import (
	"Picasso/internal/core/domain/artist"
	"Picasso/internal/core/ports"
	"github.com/mercadolibre/fury_go-core/pkg/web"
	"net/http"
)

type Handler struct {
	imageGeneratorService ports.ImageGeneratorService
	transcriberService    ports.TranscriberService
}

func NewHandler(imageGeneratorService ports.ImageGeneratorService, transcriberService ports.TranscriberService) *Handler {
	return &Handler{
		imageGeneratorService: imageGeneratorService,
		transcriberService:    transcriberService,
	}
}

func (handler Handler) Paint(w http.ResponseWriter, r *http.Request) error {
	c := r.Context()

	transcription, err := handler.transcriberService.Transcribe(c, []byte{})
	if err != nil {

	}

	text := transcription.Text
	image, err := handler.imageGeneratorService.Generate(c, text)
	if err != nil {

	}

	response := artist.Painting{Data: image}
	return web.EncodeJSON(w, response, http.StatusCreated)
}
