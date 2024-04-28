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
	ULID string
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
			ULID: ulid.String(),
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

	// varchar(255)の制約を満たすために255文字以内に切り捨てる
	if len(name) > 255 {
		return ErrEmptyName
	}

	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulid, err := ulid.New(ms, entropy)

	if err != nil {
		return err
	}

	userDetail := &UserDetail{
		ULID: ulid.String(),
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

	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulid, err := ulid.New(ms, entropy)

	if err != nil {
		return err
	}

	userDetail := &UserDetail{
		ULID: ulid.String(),
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
