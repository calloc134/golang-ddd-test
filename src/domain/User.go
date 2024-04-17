package domain

import "errors"

// UserAggregate is the aggregate root.
type UserAggregate struct {
	Name string
	Age  int
}

func NewUserAggregate(name string, age int) (*UserAggregate, error) {

	if name == "" {
		return nil, ErrEmptyName
	}

	return &UserAggregate{
		Name: name,
		Age:  age,
	}, nil
}

func (u *UserAggregate) GetName() string {
	return u.Name
}

func (u *UserAggregate) GetAge() int {
	return u.Age
}

func (u *UserAggregate) SetName(name string) error {
	if name == "" {
		return ErrEmptyName
	}

	u.Name = name
	return nil
}

func (u *UserAggregate) SetAge(age int) error {
	if 0 < age && age < 90 {
		return ErrInvalidAge
	}

	u.Age = age
	return nil
}

var (
	ErrEmptyName  = errors.New("empty name")
	ErrInvalidAge = errors.New("invalid age")
)