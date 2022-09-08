package employee

import (
	"context"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
)

var _ Service = &service{}

type Service interface {
	GetAll(ctx context.Context) ([]Employee, error)
	GetById(ctx context.Context, id int) (Employee, error)
	Create(ctx context.Context, e CreateEmployeeDTO) (int, error)
	Update(ctx context.Context, e UpdateEmployeeDTO) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
	logger     *logging.Logger
}

func (s service) GetAll(ctx context.Context) ([]Employee, error) {
	return s.repository.FindAll(ctx)
}

func (s service) GetById(ctx context.Context, id int) (Employee, error) {
	return s.repository.FindById(ctx, id)
}

func (s service) Create(ctx context.Context, e CreateEmployeeDTO) (int, error) {
	return s.repository.Create(ctx, e)
}

func (s service) Update(ctx context.Context, e UpdateEmployeeDTO) error {
	return s.repository.Update(ctx, e)
}

func (s service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func NewService(repository Repository, logger *logging.Logger) Service {
	return &service{logger: logger, repository: repository}
}
