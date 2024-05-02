package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func GenerateULID() (string, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	ulid, err := ulid.New(ms, entropy)

	if err != nil {
		return "", err
	}

	return ulid.String(), nil
}
