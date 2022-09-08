package employee

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Employee, error)
	FindById(ctx context.Context, id int) (Employee, error)
	Create(ctx context.Context, e CreateEmployeeDTO) (int, error)
	Update(ctx context.Context, e UpdateEmployeeDTO) error
	Delete(ctx context.Context, id int) error
}
