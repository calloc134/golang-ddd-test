package application

import (
	"context"

	"github.com/calloc134/golang-ddd-test/src/domain"
)

// UserAggregate is a struct that represents the User aggregate
type IUserRepository interface {
	FindAll(context context.Context) ([]domain.UserAggregate, error)
	FindByUlid(context context.Context, uuid string) (domain.UserAggregate, error)
	Save(context context.Context, user domain.UserAggregate) error
}

type UserApplication struct {
	UserRepository IUserRepository
}

func NewUserApplication(userRepository IUserRepository) UserApplication {
	return UserApplication{UserRepository: userRepository}
}

func (ua *UserApplication) FindAll(context context.Context) ([]domain.UserAggregate, error) {
	return ua.UserRepository.FindAll(context)
}

func (ua *UserApplication) FindByUlid(context context.Context, uuid string) (domain.UserAggregate, error) {
	return ua.UserRepository.FindByUlid(context, uuid)
}

func (ua *UserApplication) Save(context context.Context, user domain.UserAggregate) error {
	return ua.UserRepository.Save(context, user)
}
