package user

import "context"

type Repository interface {
	FindAll(ctx context.Context) (u []User, err error)
}
