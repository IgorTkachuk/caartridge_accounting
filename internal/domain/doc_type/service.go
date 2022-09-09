package doc_type

import (
	"context"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
)

var _ Service = &service{}

type Service interface {
	FindAll(ctx context.Context) ([]DocType, error)
	FindById(ctx context.Context, id int) (DocType, error)
	Create(ctx context.Context, docType CreateDocTypeDTO) (int, error)
	Update(ctx context.Context, docType UpdateDocTypeDTO) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
	logger     *logging.Logger
}

func (s service) FindAll(ctx context.Context) ([]DocType, error) {
	return s.repository.GetAll(ctx)
}

func (s service) FindById(ctx context.Context, id int) (DocType, error) {
	return s.repository.GetById(ctx, id)
}

func (s service) Create(ctx context.Context, docType CreateDocTypeDTO) (int, error) {
	return s.repository.Create(ctx, docType)
}

func (s service) Update(ctx context.Context, docType UpdateDocTypeDTO) error {
	return s.repository.Update(ctx, docType)
}

func (s service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func NewService(repository Repository, logger *logging.Logger) Service {
	return &service{repository: repository, logger: logger}
}
