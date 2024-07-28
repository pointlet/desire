package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
			username := c.FormValue("username")
			password := c.FormValue("password")

			fmt.Printf("User entered username: %s\nUser entered password: %s\n", username, password)

			userAccount, err := userRepository.GetUserAccountEntry(username)
			if err != nil {
				return c.JSON(401, map[string]string{"error": "Unauthorized"})
			}

			if !IsUserAuthenticated(userAccount.PasswordHash, password) {
				return c.JSON(401, map[string]string{"error": "Unauthorized"})
			}

			token, err := GenerateJWT(strconv.Itoa(userAccount.ID))
			if err != nil {
				return c.JSON(500, map[string]string{"error": "Failed to generate token"})
			}

			jwtCookie := new(http.Cookie)
			jwtCookie.Name = "token"
			jwtCookie.Value = token
			jwtCookie.Expires = time.Now().Add(72 * time.Hour)
			jwtCookie.HttpOnly = true
			jwtCookie.Secure = true
			c.SetCookie(jwtCookie)

			return next(c)
		}
	}
}
