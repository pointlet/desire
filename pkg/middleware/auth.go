package middleware

import (
	"fmt"

	"github.com/feldtsen/farrago/pkg/db"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func IsUserAuthenticated(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+db.GetPepper()))
	return err == nil
}

func Authenticate(userRepository db.UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRepository.GetUserAccountEntry("test")

			username := c.Request().Header.Get("username")
			password := c.Request().Header.Get("password")

			userAccount, err := userRepository.GetUserAccountEntry(username)
			if err != nil {
				return c.JSON(401, map[string]string{"error": "Failed"})
			}

			fmt.Printf("User entered username: %s\nUser entered password: %s\n", username, password)

			if !IsUserAuthenticated(userAccount.PasswordHash, password) {
				return c.JSON(401, map[string]string{"error": "Unauthorized"})
			}

			return next(c)
		}
	}
}
