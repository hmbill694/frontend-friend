package models

import (
	"context"
	"os"

	"github.com/tmc/langchaingo/llms/googleai"
)

type GeminiModel struct {
	version string
	llm     *googleai.GoogleAI
}

func NewGeminiModel(version string) (*GeminiModel, error) {
	apiKey := os.Getenv("GOOGLE_API_KEY")

	llm, err := googleai.New(context.Background(), googleai.WithAPIKey(apiKey), func(o *googleai.Options) {
		o.DefaultModel = version
	})

	if err != nil {
		return nil, err
	}

	return &GeminiModel{
		version,
		llm,
	}, nil
}

func (model *GeminiModel) Call(ctx context.Context, prompt string) (string, error) {
	response, err := model.llm.Call(ctx, prompt)

	if err != nil {
		return "", err
	}

	return response, nil
}
