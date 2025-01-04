package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPassword generates a bcrypt hashed password from the plain text
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	return string(hash)
}

// CheckPasswordHash compares a hashed password with its plain text equivalent
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
