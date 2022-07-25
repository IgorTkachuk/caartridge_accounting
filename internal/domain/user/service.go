package user

import (
	"context"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"log"
)

type service struct {
	repository Repository
	logger     *logging.Logger
}

type Service interface {
	GetAll(ctx context.Context) (u []User, err error)
	Create(ctx context.Context, u CreateUserDTO) (usr User, err error)
	GetById(ctx context.Context, id int) (User, error)
	GetByName(ctx context.Context, name string) (User, error)
}

func NewService(repository Repository, logger *logging.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s service) GetAll(ctx context.Context) (u []User, err error) {
	log.Println("Get all users")
	u, err = s.repository.FindAll(ctx)

	return
}

func (s service) GetById(ctx context.Context, id int) (User, error) {
	s.logger.Trace(fmt.Sprintf("Get user by ID: %d", id))
	u, err := s.repository.FindById(ctx, id)
	if err != nil {
		return User{}, err
	}

	return u, err
}

func (s service) GetByName(ctx context.Context, name string) (User, error) {
	s.logger.Trace(fmt.Sprintf("Get user by Name: %s", name))
	u, err := s.repository.FindByName(ctx, name)
	if err != nil {
		return User{}, err
	}

	return u, err
}

func (s service) Create(ctx context.Context, dto CreateUserDTO) (User, error) {
	s.logger.Debug("check password and repeat password")
	if dto.Password != dto.RepeatPassword {
		return User{}, apperror.BadRequestError("password does not match repeat password")
	}

	user := NewUser(dto)
	s.logger.Debug("generate password hash")
	err := user.GeneratePasswordHash()
	if err != nil {
		s.logger.Errorf("failed to create user due to error %v", err)
		return User{}, err
	}

	userId, err := s.repository.Create(ctx, user)
	if err != nil {
		return User{}, err
	}

	user.ID = userId

	return user, nil
}
