package user

import (
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) (u []User, err error)
	Create(ctx context.Context, u User) (int, error)
	FindById(ctx context.Context, id int) (User, error)
	FindByName(ctx context.Context, name string) (User, error)
}
