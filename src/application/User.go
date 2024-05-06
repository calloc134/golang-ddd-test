package application

import (
	"context"

	"github.com/calloc134/golang-ddd-test/src/domain"
)

// UserAggregate is a struct that represents the User aggregate
type IUserRepository interface {
	FindAll(context context.Context) ([]domain.User, error)
	FindByUlid(context context.Context, ulid domain.UlidValue) (*domain.User, error)
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

func (ua *UserApplication) NewUser(context context.Context, nameString string, ageInt int) (*domain.User, error) {

	name, err := domain.NewName(nameString)

	if err != nil {
		return nil, err
	}

	age, err := domain.NewAge(ageInt)

	if err != nil {
		return nil, err
	}

	user, err := domain.NewUser(name, age)

	if err != nil {
		return nil, err
	}

	return ua.UserRepository.Save(context, user)
}

func (ua *UserApplication) FindByUlid(context context.Context, ulidString string) (*domain.User, error) {

	ulid, err := domain.NewULID(ulidString)

	if err != nil {
		return nil, err
	}

	return ua.UserRepository.FindByUlid(context, ulid)
}

func (ua *UserApplication) UpdateNameByUlid(context context.Context, ulidString string, nameString string) (*domain.User, error) {

	name, err := domain.NewName(nameString)

	if err != nil {
		return nil, err
	}

	ulid, err := domain.NewULID(ulidString)

	if err != nil {
		return nil, err
	}

	user, err := ua.UserRepository.FindByUlid(context, ulid)

	if err != nil {
		return nil, err
	}

	if err := user.SetName(name); err != nil {
		return nil, err
	}

	return ua.UserRepository.Save(context, user)
}

func (ua *UserApplication) UpdateAgeByUlid(context context.Context, ulidString string, ageInt int) (*domain.User, error) {

	age, err := domain.NewAge(ageInt)

	if err != nil {
		return nil, err
	}

	ulid, err := domain.NewULID(ulidString)

	if err != nil {
		return nil, err
	}

	user, err := ua.UserRepository.FindByUlid(context, ulid)

	if err != nil {
		return nil, err
	}

	if err := user.SetAge(age); err != nil {
		return nil, err
	}

	return ua.UserRepository.Save(context, user)
}
