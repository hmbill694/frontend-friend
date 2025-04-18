package models

import "context"

type BaseChatModel interface {
	Call(ctx context.Context, prompt string) (string, error)
}
