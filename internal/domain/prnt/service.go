package prnt

import (
	"context"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
)

var _ Service = &service{}

type Service interface {
	GetAll(ctx context.Context) (p []Prn, err error)
	Create(ctx context.Context, p CreatePrnDTO) (int, error)
	GetById(ctx context.Context, id int) (v Prn, err error)
	GetByName(ctx context.Context, name string) ([]Prn, error)
	Delete(ctx context.Context, id int) (int, error)
	Update(ctx context.Context, p UpdatePrnDTO) error
}

type service struct {
	repository Repository
	logger     *logging.Logger
}

func (s service) GetAll(ctx context.Context) (p []Prn, err error) {
	s.logger.Trace("Get all printer model records")
	return s.repository.FindAll(ctx)
}

func (s service) Create(ctx context.Context, p CreatePrnDTO) (int, error) {
	return s.repository.Create(ctx, p)
}

func (s service) GetById(ctx context.Context, id int) (v Prn, err error) {
	return s.repository.FindById(ctx, id)
}

func (s service) GetByName(ctx context.Context, name string) ([]Prn, error) {
	return s.repository.FindByName(ctx, name)
}

func (s service) Delete(ctx context.Context, id int) (int, error) {
	return s.repository.Delete(ctx, id)
}

func (s service) Update(ctx context.Context, p UpdatePrnDTO) error {
	return s.repository.Update(ctx, p)
}

func NewService(repository Repository, logger *logging.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}
