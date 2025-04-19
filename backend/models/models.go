package models

import "context"

type ChatModel interface {
	Call(ctx context.Context, prompt string) (string, error)
}
