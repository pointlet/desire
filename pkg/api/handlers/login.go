package handlers

import (
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
	csrfToken := GetCSRFToken(c)
	return renderByHXRequest(c, login.LoginPartial(csrfToken), login.LoginPage(csrfToken))
}

func (h *Login) PostHandler(c echo.Context) error {
	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Redirect", "/home")
		return c.NoContent(http.StatusOK)
	}

	return c.Redirect(http.StatusFound, "/home")
}
