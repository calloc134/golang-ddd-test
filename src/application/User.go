package application

import (
	"context"

	"github.com/calloc134/golang-ddd-test/src/domain"
)

// UserAggregate is a struct that represents the User aggregate
type IUserRepository interface {
	FindAll(context context.Context) ([]domain.UserAggregate, error)
	FindByUlid(context context.Context, uuid string) (*domain.UserAggregate, error)
	Save(context context.Context, user *domain.UserAggregate) (*domain.UserAggregate, error)
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

func (ua *UserApplication) NewUser(context context.Context, name string, age int) (*domain.UserAggregate, error) {
	user, err := domain.NewUserAggregate(name, age)

	if err != nil {
		return nil, err
	}

	return ua.UserRepository.Save(context, user)
}

func (ua *UserApplication) FindByUlid(context context.Context, uuid string) (*domain.UserAggregate, error) {
	return ua.UserRepository.FindByUlid(context, uuid)
}

func (ua *UserApplication) UpdateNameByUlid(context context.Context, uuid string, name string) (*domain.UserAggregate, error) {
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

func (ua *UserApplication) UpdateAgeByUlid(context context.Context, uuid string, age int) (*domain.UserAggregate, error) {
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