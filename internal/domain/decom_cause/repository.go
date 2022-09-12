package decom_cause

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]DecomCause, error)
	FindById(ctx context.Context, id int) (DecomCause, error)
	Create(ctx context.Context, dto CreateDecomCauseDTO) (int, error)
	Update(ctx context.Context, dto UpdateDecomCauseDTO) error
	Delete(ctx context.Context, id int) error
}
