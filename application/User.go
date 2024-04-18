package application

import (
	"github.com/calloc134/golang-ddd-test/domain"
)

// UserAggregate is a struct that represents the User aggregate
type IUserAggregate interface {
	FindAll() ([]domain.UserAggregate, error)
	FindByID(id int) (domain.UserAggregate, error)
	Save(user domain.UserAggregate) error
}

type UserApplication struct {
	UserAggregate domain.UserAggregate
}

func NewUserApplication(userAggregate domain.UserAggregate) UserApplication {
	return UserApplication{UserAggregate: userAggregate}
}