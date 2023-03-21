package Picasso

import (
	"OpenAIDemo/cmd/dependencies"
	"OpenAIDemo/cmd/httpserver"
)

func main() {
	d := dependencies.NewByEnvironment()
	httpserver.Start(d)
}
