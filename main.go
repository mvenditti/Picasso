package main

import (
	"Picasso/cmd/dependencies"
	"Picasso/cmd/httpserver"
)

func main() {
	d := dependencies.NewByEnvironment()
	httpserver.Start(d)
}
