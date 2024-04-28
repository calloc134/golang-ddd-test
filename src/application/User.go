package application

import (
	"context"

	"github.com/calloc134/golang-ddd-test/src/domain"
)

// UserAggregate is a struct that represents the User aggregate
type IUserRepository interface {
	FindAll(context context.Context) ([]domain.User, error)
	FindByUlid(context context.Context, uuid string) (*domain.User, error)
	Save(context context.Context, user *domain.User) (*domain.User, error)
}

type UserApplication struct {
	UserRepository IUserRepository
}

func NewUserApplication(userRepository IUserRepository) UserApplication {
	return UserApplication{UserRepository: userRepository}
}

func (ua *UserApplication) FindAll(context context.Context) ([]domain.User, error) {
	return ua.UserRepository.FindAll(context)
}

func (ua *UserApplication) NewUser(context context.Context, name string, age int) (*domain.User, error) {
	user, err := domain.NewUser(name, age)

	if err != nil {
		return nil, err
	}

	return ua.UserRepository.Save(context, user)
}

func (ua *UserApplication) FindByUlid(context context.Context, uuid string) (*domain.User, error) {
	return ua.UserRepository.FindByUlid(context, uuid)
}

func (ua *UserApplication) UpdateNameByUlid(context context.Context, uuid string, name string) (*domain.User, error) {
	user, err := ua.FindByUlid(context, uuid)

	if err != nil {
		return nil, err
	}

	err = user.SetName(name)

	if err != nil {
		return nil, err
	}

	return ua.UserRepository.Save(context, user)
}

func (ua *UserApplication) UpdateAgeByUlid(context context.Context, uuid string, age int) (*domain.User, error) {
	user, err := ua.FindByUlid(context, uuid)

	if err != nil {
		return nil, err
	}

	err = user.SetAge(age)

	if err != nil {
		return nil, err
	}

	return ua.UserRepository.Save(context, user)
}
