package ou

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Ou, error)
	FindById(ctx context.Context, id int) (Ou, error)
	Create(ctx context.Context, ou CreateOuDTO) (int, error)
	Update(ctx context.Context, ou UpdateOuDTO) error
	Delete(ctx context.Context, id int) error
}
