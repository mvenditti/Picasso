package httpserver

import (
	"Picasso/cmd/dependencies"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

func routes(d dependencies.Definition) *fury.Application {
	router, err := fury.NewWebApplication()

	router.Get("/image", d.ArtistHandler.Paint)

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
