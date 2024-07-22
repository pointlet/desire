package middleware

import (
	"errors"
	"strings"

	"github.com/feldtsen/farrago/pkg/db"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var PEPPER = "d0924ce78d7279c0f694799a" // TODO: remove this and set as env var

func CreateNewUser(statements *db.PreparedStatements, username, password string) error {
	hashedPassword, err := GenerateHashPassword(password)

	if err != nil {
		return err
	}

	_, err = statements.InsertNewUser.Exec(username, hashedPassword)

	if err != nil {
		return err
	}

	return nil

}

func GenerateHashPassword(password string) (string, error) {
	if len(password) < 12 || len(password) > 40 {
		return "", errors.New("password length must be between 8 and 40 characters")
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

func IsUserAuthenticated(username, password string) bool {
	// var storedHash string

	// err := statements.GetUserAccountPasswordHash.QueryRow(username).Scan(&storedHash)

	// if err != nil {
	// 	return false
	// }

	// err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password+PEPPER))

	// return err == nil
	return true
}

func Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			username := c.Request().Header.Get("username")
			password := c.Request().Header.Get("password")

			if !IsUserAuthenticated(username, password) {
				return c.JSON(401, map[string]string{"error": "Unauthorized"})
			}

			return next(c)
		}
	}
}
