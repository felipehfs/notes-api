package infra

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorInvalidPassword = errors.New("please verify the correct fields")
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
