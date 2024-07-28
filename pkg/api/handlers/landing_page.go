package handlers

import (
	"net/http"

	"github.com/feldtsen/farrago/pkg/middleware"
	"github.com/feldtsen/farrago/view/page/landing"
	"github.com/labstack/echo/v4"
)

type LandingPage struct{}

func (h *LandingPage) Handler(c echo.Context) error {
	if !middleware.IsLoggedIn(c) {
		return renderByHXRequest(c, landing.LandingPartial(), landing.LandingPage())
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Redirect", "/home")
		return c.NoContent(http.StatusOK)
	}

	return c.Redirect(http.StatusFound, "/home")
}
