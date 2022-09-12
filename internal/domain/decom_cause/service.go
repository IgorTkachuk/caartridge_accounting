package decom_cause

import (
	"context"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
)

var _ Service = &service{}

type Service interface {
	GetAll(ctx context.Context) ([]DecomCause, error)
	GetById(ctx context.Context, id int) (DecomCause, error)
	Create(ctx context.Context, dto CreateDecomCauseDTO) (int, error)
	Update(ctx context.Context, dto UpdateDecomCauseDTO) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
	logger     *logging.Logger
}

func (s service) GetAll(ctx context.Context) ([]DecomCause, error) {
	return s.repository.FindAll(ctx)
}

func (s service) GetById(ctx context.Context, id int) (DecomCause, error) {
	return s.repository.FindById(ctx, id)
}

func (s service) Create(ctx context.Context, dto CreateDecomCauseDTO) (int, error) {
	return s.repository.Create(ctx, dto)
}

func (s service) Update(ctx context.Context, dto UpdateDecomCauseDTO) error {
	return s.repository.Update(ctx, dto)
}

func (s service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func NewService(repository Repository, logger *logging.Logger) Service {
	return &service{repository: repository, logger: logger}
}
