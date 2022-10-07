package ctr_showcase

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]CtrShowcaseDTO, error)
}
