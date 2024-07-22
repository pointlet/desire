package handlers

import (
	"net/http"

	"github.com/feldtsen/farrago/view/page/login"
	"github.com/labstack/echo/v4"
)

type Login struct{}

func (h *Login) GetHandler(c echo.Context) error {
	return renderByHXRequest(c, login.LoginPartial(), login.LoginPage())
}

func (h *Login) PostHandler(c echo.Context) error {
	print("login clicked")
	return c.String(http.StatusOK, "login clicked successfully, bitch")
}
