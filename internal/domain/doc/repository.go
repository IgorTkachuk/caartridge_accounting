package doc

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Doc, error)
	FindById(ctx context.Context, id int) (Doc, error)
	Create(ctx context.Context, dto CreateDocDTO) (int, error)
	Update(ctx context.Context, dto UpdateDocDTO) error
	Delete(ctx context.Context, id int) error
}
