package artist

import "github.com/sashabaranov/go-openai"

type Painting struct {
	Data openai.ImageResponse `json:"data"`
}
