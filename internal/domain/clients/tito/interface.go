package tito

import (
	"context"
)

type ITitoClient interface {
	SendMessage(ctx context.Context, message string) (interface{}, error)
}
