package handlers

import (
	"fmt"
	"net/http"

	"github.com/feldtsen/farrago/view/page/login"
	"github.com/labstack/echo/v4"
)

type UserAccountRequset struct {
	Username string
	Password string
}

type Login struct{}

func (h *Login) GetHandler(c echo.Context) error {
	return renderByHXRequest(c, login.LoginPartial(), login.LoginPage())
}

func (h *Login) PostHandler(c echo.Context) error {
	fmt.Println("this is after authentication")

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Redirect", "/home")
		return c.NoContent(http.StatusOK)
	}

	return c.Redirect(http.StatusFound, "/home")
}
