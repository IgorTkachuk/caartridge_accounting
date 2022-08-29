package prnt

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Prn, error)
	Create(ctx context.Context, prn CreatePrnDTO) (id int, err error)
	FindById(ctx context.Context, id int) (Prn, error)
	FindByName(ctx context.Context, name string) ([]Prn, error)
	Delete(ctx context.Context, id int) (int, error)
	Update(ctx context.Context, prn UpdatePrnDTO) error
}
