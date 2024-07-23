package db

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var PEPPER = "d0924ce78d7279c0f694799a" // TODO: remove this and set new pepper as env var

// TODO: REMOVE THIS
func GetPepper() string {
	return PEPPER
}

func GenerateHashPassword(password string) (string, error) {
	if len(password) < 12 || len(password) > 40 {
		return "", errors.New("password length must be between 12 and 40 characters")
	}

	if strings.Contains(password, " ") {
		return "", errors.New("password cannot contain white space")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+PEPPER), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
