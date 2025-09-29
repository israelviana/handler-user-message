package repository

import "context"

type IRepository interface {
	SaveIncomingMessage(ctx context.Context, data string) error
}
