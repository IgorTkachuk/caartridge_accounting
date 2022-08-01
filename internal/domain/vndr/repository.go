package vndr

import "context"

type Repository interface {
	FindAll(ctx context.Context) (v []Vendor, err error)
	Create(ctx context.Context, v Vendor) (int, error)
	FindById(ctx context.Context, id int) (v Vendor, err error)
	FindByName(ctx context.Context, name string) ([]Vendor, error)
}
