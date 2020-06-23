package account

import "context"

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, string)
	GetUser(ctx context.Context, id string) (string, error)
}
