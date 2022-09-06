package business_line

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]BusinessLine, error)
	FindById(ctx context.Context, id int) (BusinessLine, error)
	Create(ctx context.Context, bl CreateBusinessLineDTO) (int, error)
	Update(ctx context.Context, bl UpdateBusinessLineDTO) error
	Delete(ctx context.Context, id int) error
}
