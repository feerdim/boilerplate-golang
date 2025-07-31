package util

import (
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (hashed string, err error) {
	costStr := os.Getenv("AUTH_BCRYPT_COST")

	cost, err := strconv.Atoi(costStr)
	if err != nil {
		err = errors.Wrapf(err, "error parse int on bcrypt cost env : %s", costStr)
		return
	}

	if cost < bcrypt.MinCost {
		cost = bcrypt.DefaultCost
	}

	crypt, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		err = errors.Wrap(err, "error generate bcrypt hash password")
		return
	}

	hashed = string(crypt)

	return
}

func CompareHashPassword(passwordInput, passwordDB string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordInput))
	return
}

func GenerateUUID() string {
	guid, err := uuid.NewV7()
	if err != nil {
		return uuid.NewString()
	}

	return guid.String()
}
