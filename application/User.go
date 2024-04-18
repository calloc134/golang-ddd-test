package application

import (
	"github.com/calloc134/golang-ddd-test/domain"
)

// UserAggregate is a struct that represents the User aggregate
type IUserRepository interface {
	FindAll() ([]domain.UserAggregate, error)
	FindByID(id int) (domain.UserAggregate, error)
	Save(user domain.UserAggregate) error
}

type UserApplication struct {
	UserRepository IUserRepository
}

func NewUserApplication(userRepository IUserRepository) UserApplication {
	return UserApplication{UserRepository: userRepository}
}

func (ua *UserApplication) FindAll() ([]domain.UserAggregate, error) {
	return ua.UserRepository.FindAll()
}

func (ua *UserApplication) FindByID(id int) (domain.UserAggregate, error) {
	return ua.UserRepository.FindByID(id)
}

