package cartridge_model

import (
	"context"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
)

var _ Service = &service{}

type Service interface {
	GetAll(ctx context.Context) ([]CartridgeModel, error)
	GetById(ctx context.Context, id int) (CartridgeModel, error)
	Create(ctx context.Context, c CreateCartridgeModelDTO) (int, error)
	Update(ctx context.Context, c UpdateCartridgeModelDTO) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	logger     *logging.Logger
	repository Repository
}

func NewService(repository Repository, logger *logging.Logger) *service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s service) GetAll(ctx context.Context) ([]CartridgeModel, error) {
	return s.repository.FindAll(ctx)
}

func (s service) GetById(ctx context.Context, id int) (CartridgeModel, error) {
	return s.repository.FindById(ctx, id)
}

func (s service) Create(ctx context.Context, c CreateCartridgeModelDTO) (int, error) {
	return s.repository.Create(ctx, c)
}

func (s service) Update(ctx context.Context, c UpdateCartridgeModelDTO) error {
	return s.repository.Update(ctx, c)
}

func (s service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
