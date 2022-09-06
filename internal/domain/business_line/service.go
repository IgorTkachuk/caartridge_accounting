package business_line

import (
	"context"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
)

var _ Service = &service{}

type Service interface {
	GetAll(ctx context.Context) ([]BusinessLine, error)
	GetById(ctx context.Context, id int) (BusinessLine, error)
	Create(ctx context.Context, bl CreateBusinessLineDTO) (int, error)
	Update(ctx context.Context, bl UpdateBusinessLineDTO) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
	logger     *logging.Logger
}

func (s service) GetAll(ctx context.Context) ([]BusinessLine, error) {
	return s.repository.FindAll(ctx)
}

func (s service) GetById(ctx context.Context, id int) (BusinessLine, error) {
	return s.repository.FindById(ctx, id)
}

func (s service) Create(ctx context.Context, bl CreateBusinessLineDTO) (int, error) {
	return s.repository.Create(ctx, bl)
}

func (s service) Update(ctx context.Context, bl UpdateBusinessLineDTO) error {
	return s.repository.Update(ctx, bl)
}

func (s service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func NewService(repository Repository, logger *logging.Logger) *service {
	return &service{repository: repository, logger: logger}
}
