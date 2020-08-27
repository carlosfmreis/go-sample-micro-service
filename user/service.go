package user

import (
	"context"
)

// Service structure
type service struct {
	repository Repository
}

// Service interface
type Service interface {
	CreateUser(context context.Context, email string, password string) error
	GetUser(context context.Context, id int) (User, error)
}

func (service *service) CreateUser(context context.Context, email string, password string) error {
	user := User{
		Email:    email,
		Password: password,
	}
	err := service.repository.CreateUser(context, user)
	return err
}

func (service *service) GetUser(context context.Context, id int) (User, error) {
	user, err := service.repository.GetUser(context, id)
	return user, err
}

// NewService method
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
