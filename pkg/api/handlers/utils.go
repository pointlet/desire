package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func renderByHXRequest(c echo.Context, partial templ.Component, page templ.Component) error {
	if c.Request().Header.Get("HX-Request") == "true" {
		return render(c, partial)
	}

	return render(c, page)
}

func GetCSRFToken(c echo.Context) string {
	return c.Get(echoMiddleware.DefaultCSRFConfig.ContextKey).(string)
}
