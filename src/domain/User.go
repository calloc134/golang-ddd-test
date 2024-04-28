package domain

import (
	"errors"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// User is the aggregate root.
type User struct {
	ULID       string
	UserDetail *UserDetail
	Version    int
}

type UserDetail struct {
	Name string
	Age  int
}

func NewUser(name string, age int) (*User, error) {

	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulid, err := ulid.New(ms, entropy)

	if err != nil {
		return nil, err
	}

	user := &User{
		ULID:    ulid.String(),
		Version: 0,
		UserDetail: &UserDetail{
			Name: "",
			Age:  0,
		},
	}

	err = user.SetName(name)
	if err != nil {
		return nil, err
	}

	err = user.SetAge(age)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) GetName() string {
	return u.UserDetail.Name
}

func (u *User) GetAge() int {
	return u.UserDetail.Age
}

func (u *User) SetName(name string) error {
	if name == "" {
		return ErrEmptyName
	}

	u.UserDetail.Name = name
	return nil
}

func (u *User) SetAge(age int) error {
	if 0 < age && age < 90 {
		return ErrInvalidAge
	}

	u.UserDetail.Age = age
	return nil
}

var (
	ErrEmptyName  = errors.New("empty name")
	ErrInvalidAge = errors.New("invalid age")
)