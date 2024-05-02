package domain

import (
	"errors"

	"github.com/calloc134/golang-ddd-test/src/utils"
)

// User is the aggregate root.
type User struct {
	ULID       string
	UserDetail *UserDetail
	Version    int
}

type UserDetail struct {
	ULID string
	Name string
	Age  int
}

func NewUser(name string, age int) (*User, error) {

	ulid, err := utils.GenerateULID()

	if err != nil {
		return nil, err
	}

	user := &User{
		ULID:    ulid,
		Version: 0,
		UserDetail: &UserDetail{
			Name: "",
			Age:  0,
		},
	}

	if err := user.SetName(name); err != nil {
		return nil, err
	}

	if err := user.SetAge(age); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) SetName(name string) error {
	if name == "" {
		return ErrEmptyName
	}

	// varchar(255)の制約を満たすために255文字以内に切り捨てる
	if len(name) > 255 {
		return ErrEmptyName
	}

	ulid, err := utils.GenerateULID()

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

func (u *User) SetAge(age int) error {
	if age < 0 || age > 120 {
		return ErrInvalidAge
	}

	ulid, err := utils.GenerateULID()

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
