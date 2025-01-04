package utils

import (

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hashed password from the plain text
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "[HashPassword]")
	}
	return string(hash),err
}

// CheckPasswordHash compares a hashed password with its plain text equivalent
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
