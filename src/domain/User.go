package domain

import (
	"errors"
)

// User is the aggregate root.
type User struct {
	ULID       UlidValue
	UserDetail *UserDetail
	Version    int
}

type UserDetail struct {
	ULID UlidValue
	Name NameValue
	Age  AgeValue
}

type NameValue struct {
	value string
}

type AgeValue struct {
	value int
}

func (n NameValue) String() string {
	return n.value
}

func (a AgeValue) Int() int {
	return a.value
}

func NewName(value string) (NameValue, error) {
	if value == "" {
		return NameValue{}, ErrEmptyName
	}

	if len(value) > 255 {
		return NameValue{}, ErrInvalidName
	}

	return NameValue{value: value}, nil
}

func NewAge(value int) (AgeValue, error) {
	if value < 0 || value > 120 {
		return AgeValue{}, ErrInvalidAge
	}

	return AgeValue{value: value}, nil
}

func NewUser(nameString string, ageInt int) (*User, error) {

	name, err := NewName(nameString)

	if err != nil {
		return nil, err
	}

	ulid, err := GenerateULID()

	if err != nil {
		return nil, err
	}

	age, err := NewAge(ageInt)

	if err != nil {
		return nil, err
	}

	user := &User{
		ULID:    ulid,
		Version: 0,
		UserDetail: &UserDetail{
			Name: name,
			Age:  age,
		},
	}

	return user, nil
}

func (u *User) SetName(nameString string) error {

	name, err := NewName(nameString)

	if err != nil {
		return err
	}

	ulid, err := GenerateULID()

	if err != nil {
		return err
	}

	userDetail := &UserDetail{
		ULID: ulid,
		Name: name,
		Age:  u.UserDetail.Age,
	}

	u.UserDetail = userDetail
	return nil
}

func (u *User) SetAge(ageInt int) error {

	ulid, err := GenerateULID()

	if err != nil {
		return err
	}

	age, err := NewAge(ageInt)

	if err != nil {
		return err
	}

	userDetail := &UserDetail{
		ULID: ulid,
		Name: u.UserDetail.Name,
		Age:  age,
	}

	u.UserDetail = userDetail
	return nil
}

var (
	ErrEmptyName   = errors.New("empty name")
	ErrInvalidName = errors.New("invalid name")
	ErrInvalidAge  = errors.New("invalid age")
)
