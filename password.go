package main

import (
	"fmt"

	vsLog "github.com/dhf0820/vslog"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	vsLog.Info("hashedPassword: " + string(hashedPassword))
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	vsLog.Debug2(fmt.Sprintf("Checking Password: %s against: %s", password, hashedPassword))
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
