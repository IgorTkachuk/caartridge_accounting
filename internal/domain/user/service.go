package user

import (
	"context"
	"log"
)

type service struct {
	repository Repository
}

type Service interface {
	GetAll(ctx context.Context) (u []User, err error)
	Create(ctx context.Context, u CreateUserDTO) (id int, err error)
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) GetAll(ctx context.Context) (u []User, err error) {
	log.Println("Get all users")
	u, err = s.repository.FindAll(ctx)

	return
}

func (service) Create(ctx context.Context, u CreateUserDTO) (id int, err error) {
	// TODO метод репозитория подготовлен. Нужно реадизовать сервис
	panic("Implement me")
}
