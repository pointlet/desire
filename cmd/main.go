package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/feldtsen/farrago/pkg/api/handlers"
	"github.com/feldtsen/farrago/pkg/config"
	"github.com/feldtsen/farrago/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	// Setup
	e := initEchoServer()
	defer e.Close()

	dbpool, err := config.ConnectToDB()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	defer dbpool.Close()

	// Landing page
	landingPage := &handlers.LandingPage{}
	e.GET("/", landingPage.Handler)

	// Desire path
	desirePath := &handlers.DesirePath{}
	e.GET("/desirePath", middleware.Pagination(desirePath.Handler))

	// Authentication
	login := &handlers.Login{}
	e.GET("/login", login.GetHandler)
	e.POST("/login", login.PostHandler)

	// Super secret page
	e.GET("/secret", func(c echo.Context) error {
		return c.String(http.StatusOK, "You are authenticated!")
	}, middleware.Authenticate())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

func initEchoServer() *echo.Echo {
	e := echo.New()

	// Static files
	e.Static("/static", "static")

	// Set logging level
	e.Logger.SetLevel(log.INFO)

	return e
}
