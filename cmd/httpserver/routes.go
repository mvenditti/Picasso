package httpserver

import (
	"OpenAIDemo/cmd/dependencies"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

func routes(d dependencies.Definition) *fury.Application {
	router, err := fury.NewWebApplication()

	if err != nil {
		panic(err)
	}

	return router
}

func Start(d dependencies.Definition) {
	if err := routes(d).Run(); err != nil {
		panic(err)
	}
}
