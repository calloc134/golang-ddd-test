package domain

import (
	"errors"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// UserAggregate is the aggregate root.
type UserAggregate struct {
	ULID  string
	Name string
	Age  int
}

func NewUserAggregate(name string, age int) (*UserAggregate, error) {

	if name == "" {
		return nil, ErrEmptyName
	}

	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulid, err := ulid.New(ms, entropy)

	if err != nil {
		return nil, err
	}

	return &UserAggregate{
		ULID: ulid.String(),
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