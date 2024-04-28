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

func NewUserAggregate() (*User, error) {

	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulid, err := ulid.New(ms, entropy)

	if err != nil {
		return nil, err
	}

	return &User{
		ULID:    ulid.String(),
		Name:    "",
		Age:     -1,
		Version: 0,
	}, nil
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetAge() int {
	return u.Age
}

func (u *User) SetName(name string) error {
	if name == "" {
		return ErrEmptyName
	}

	u.Name = name
	return nil
}

func (u *User) SetAge(age int) error {
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
