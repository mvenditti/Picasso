package dependencies

import (
	"Picasso/cmd/configurations"
	"Picasso/internal/core/services/image_generator"
	"Picasso/internal/core/services/transcriber"
	"Picasso/internal/handlers/artist"
	"Picasso/pkg/constants"
	"Picasso/pkg/environment"
	"strings"
)

type Definition struct {
	ArtistHandler *artist.Handler
}

func NewByEnvironment() Definition {
	d := Definition{}

	envManager := environment.New()
	env := getEnvironment(envManager)
	client := configurations.InitOpenAIClient(env)

	imageGeneratorService := image_generator.NewService(client)
	transcriberService := transcriber.NewService(client)

	d.ArtistHandler = artist.NewHandler(imageGeneratorService, transcriberService)

	return d
}

func getEnvironment(manager environment.Manager) string {
	env := manager.GetEnv(constants.GoEnvironment)
	result := constants.Local

	if env == constants.Production {
		result = constants.Production
	}

	scope := manager.GetEnv(constants.Scope)
	if strings.Contains(scope, constants.Staging) {
		result = constants.Staging
	}

	if strings.Contains(scope, constants.Sandbox) {
		result = constants.Sandbox
	}

	return result
}
