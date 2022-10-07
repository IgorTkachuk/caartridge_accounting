package ctr_showcase

import (
	"context"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
)

var _ Service = &service{}

type Service interface {
	GetAll(ctx context.Context) ([]CtrShowcaseDTO, error)
}

type service struct {
	repository Repository
	logger     *logging.Logger
}

func NewService(repository Repository, logger *logging.Logger) *service {
	return &service{repository: repository, logger: logger}
}

func (s service) GetAll(ctx context.Context) ([]CtrShowcaseDTO, error) {
	return s.repository.FindAll(ctx)
}
