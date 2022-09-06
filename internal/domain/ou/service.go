package ou

import (
	"context"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
)

var _ Service = &service{}

type Service interface {
	GetAll(ctx context.Context) ([]Ou, error)
	GetById(ctx context.Context, id int) (Ou, error)
	Create(ctx context.Context, ou CreateOuDTO) (int, error)
	Update(ctx context.Context, ou UpdateOuDTO) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
	logger     *logging.Logger
}

func NewService(repository Repository, logger *logging.Logger) *service {
	return &service{repository: repository, logger: logger}
}

func (s service) GetAll(ctx context.Context) ([]Ou, error) {
	return s.repository.FindAll(ctx)
}

func (s service) GetById(ctx context.Context, id int) (Ou, error) {
	return s.repository.FindById(ctx, id)
}

func (s service) Create(ctx context.Context, ou CreateOuDTO) (int, error) {
	return s.repository.Create(ctx, ou)
}

func (s service) Update(ctx context.Context, ou UpdateOuDTO) error {
	return s.repository.Update(ctx, ou)
}

func (s service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
