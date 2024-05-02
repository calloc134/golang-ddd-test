package domain

import (
	"errors"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type UlidValue struct {
	value string
}

func GenerateULID() (UlidValue, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulid, err := ulid.New(ms, entropy)

	if err != nil {
		return UlidValue{}, err
	}

	ulidValue, err := NewULID(ulid.String())

	if err != nil {
		return UlidValue{}, err
	}

	return ulidValue, nil
}

func NewULID(value string) (UlidValue, error) {
	if value == "" {
		return UlidValue{}, ErrEmptyULID
	}

	if len(value) != 26 {
		return UlidValue{}, ErrInvalidULID
	}

	return UlidValue{value: value}, nil
}

func (u *UlidValue) Equals(target *UlidValue) bool {
	return u.value == target.value
}

var (
	ErrEmptyULID   = errors.New("empty ulid")
	ErrInvalidULID = errors.New("invalid ulid")
)
