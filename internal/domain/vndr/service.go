package vndr

import (
	"context"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
)

var _ Service = &service{}

type Service interface {
	GetAll(ctx context.Context) (v []Vendor, err error)
	Create(ctx context.Context, v CreateVendorDTO) (int, error)
	GetById(ctx context.Context, id int) (v Vendor, err error)
	GetByName(ctx context.Context, name string) ([]Vendor, error)
}

type service struct {
	repository Repository
	logger     *logging.Logger
}

func NewService(repository Repository, logger *logging.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s service) GetAll(ctx context.Context) (v []Vendor, err error) {
	s.logger.Trace("Get all vendors records")
	v, err = s.repository.FindAll(ctx)

	return
}

func (s service) Create(ctx context.Context, v CreateVendorDTO) (id int, err error) {
	vndr := NewVendor(v)
	id, err = s.repository.Create(ctx, vndr)

	return
}

func (s service) GetById(ctx context.Context, id int) (v Vendor, err error) {
	return s.repository.FindById(ctx, id)
}

func (s service) GetByName(ctx context.Context, name string) ([]Vendor, error) {
	return s.repository.FindByName(ctx, name)
}
