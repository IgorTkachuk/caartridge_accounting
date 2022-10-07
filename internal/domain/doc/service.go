package doc

import (
	"context"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
)

var _ Service = &service{}

type Service interface {
	GetAll(ctx context.Context) ([]Doc, error)
	GetById(ctx context.Context, id int) (Doc, error)
	Create(ctx context.Context, dto CreateDocDTO) (int, error)
	Update(ctx context.Context, dto UpdateDocDTO) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
	logger     *logging.Logger
}

func (s service) GetAll(ctx context.Context) ([]Doc, error) {
	return s.repository.FindAll(ctx)
}

func (s service) GetById(ctx context.Context, id int) (Doc, error) {
	return s.repository.FindById(ctx, id)
}

func (s service) Create(ctx context.Context, dto CreateDocDTO) (int, error) {
	return s.repository.Create(ctx, dto)
}

func (s service) Update(ctx context.Context, dto UpdateDocDTO) error {
	return s.repository.Update(ctx, dto)
}

func (s service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func NewService(repository Repository, logger *logging.Logger) *service {
	return &service{repository: repository, logger: logger}
}
