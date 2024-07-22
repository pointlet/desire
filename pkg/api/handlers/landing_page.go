package handlers

import (
	"github.com/feldtsen/farrago/view/page/landing"
	"github.com/labstack/echo/v4"
)

type LandingPage struct{}

func (h *LandingPage) Handler(c echo.Context) error {
	return renderByHXRequest(c, landing.LandingPartial(), landing.LandingPage())
}
