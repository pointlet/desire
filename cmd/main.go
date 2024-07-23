package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/feldtsen/farrago/pkg/api/handlers"
	"github.com/feldtsen/farrago/pkg/config"
	"github.com/feldtsen/farrago/pkg/db"
	"github.com/feldtsen/farrago/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Server struct {
	e      *echo.Echo
	config *config.Config
}

func NewServer() *Server {
	return &Server{
		e:      echo.New(),
		config: config.NewConfig(),
	}
}

func main() {

	server := NewServer()

	server.e.Logger.SetLevel(log.INFO)
	server.e.Static("/static", "static")
	defer server.e.Close()

	dbpool, err := db.ConnectToDB(&server.config.DB)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer dbpool.Close()

	// TODO: testing interface for UserRepository
	userRepository := &db.PgxUserRepository{DB: dbpool}

	userRepository.InsertUserAccountEntry("test", "test")
	userAccount, err := userRepository.GetUserAccountEntry("test")
	if err != nil {
		log.Error(err)
	} else {
		log.Printf("User account: %v", userAccount)
	}
	userRepository.DeleteUserAccountEntry("test")

	// Landing page
	landingPage := &handlers.LandingPage{}
	server.e.GET("/", landingPage.Handler)

	// Desire path
	desirePath := &handlers.DesirePath{}
	server.e.GET("/desirePath", middleware.Pagination(desirePath.Handler))

	// Authentication
	login := &handlers.Login{}
	server.e.GET("/login", login.GetHandler)
	server.e.POST("/login", login.PostHandler)

	// Super secret page
	server.e.GET("/secret", func(c echo.Context) error {
		return c.String(http.StatusOK, "You are authenticated!")
	}, middleware.Authenticate(userRepository))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start server
	go func() {
		if err := server.e.Start(server.config.Server.Port); err != nil && err != http.ErrServerClosed {
			server.e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.e.Shutdown(ctx); err != nil {
		server.e.Logger.Fatal(err)
	}
}
