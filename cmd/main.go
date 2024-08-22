package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/feldtsen/farrago/pkg/api/handlers"
	"github.com/feldtsen/farrago/pkg/config"
	"github.com/feldtsen/farrago/pkg/db"
	"github.com/feldtsen/farrago/pkg/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
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

// TODO: remove this
func sanityTestCRUD(userRepository db.UserRepository) {

	// registration flow
	insertResult, err := userRepository.InsertUserAccountEntry("test", "1234abcd!@#$")
	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("Insert user account, rows affected: %d\n", insertResult.RowsAffected)
	}
	userAccount, err := userRepository.GetUserAccountEntry("test")
	if err == nil {
		fmt.Printf("Username: %s\nPassword: %s\n", userAccount.Username, userAccount.PasswordHash)
	}

	// is authenticated???
	intruderAccount := middleware.IsUserAuthenticated("intruder", "intruderpassword")
	legitAccount := middleware.IsUserAuthenticated(userAccount.PasswordHash, "1234abcd!@#$")
	wrongPasswordAccount := middleware.IsUserAuthenticated(userAccount.PasswordHash, "1234abcd!@#$ ")

	fmt.Printf("Intruder is authenticated: %v\nLegit user is authenticated: %v\nWrong password user is authenticated: %v\n", intruderAccount, legitAccount, wrongPasswordAccount)

	// clean up
	// userRepository.DeleteUserAccountEntry("test")
}

func main() {
	server := NewServer()

	//TODO: workaround for not caching static during dev
	// TOOO: correct solution would be to generate a hash for the static files and use that as a query parameter
	 server.e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
				if strings.HasPrefix(c.Path(), "/static") {
						c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
				}
				return next(c)
		}
})

	server.e.Static("/static", "static")

	server.e.Use(echoMiddleware.SecureWithConfig(echoMiddleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))

	// TODO: I need to make sure that we accept localhost during local developer, but not in prod
	server.e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		// AllowOrigins: []string{"http://localhost:*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	server.e.Use(echoMiddleware.CSRFWithConfig(echoMiddleware.CSRFConfig{
		TokenLength:    32,
		TokenLookup:    "form:csrf_token",
		CookieName:     "_csrf",
		CookieMaxAge:   86400,
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookieSameSite: http.SameSiteStrictMode,
	}))

	server.e.Logger.SetLevel(log.INFO)
	server.e.Use()
	defer server.e.Close()

	dbpool, err := db.ConnectToDB(&server.config.DB)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer dbpool.Close()

	userRepository := &db.PgxUserRepository{DB: dbpool}
	// TODO: testing interface for UserRepository
	sanityTestCRUD(userRepository)

	// Landing page
	landingPage := &handlers.LandingPage{}
	server.e.GET("/", landingPage.Handler)

	// Desire path
	desirePath := &handlers.DesirePath{}
	server.e.GET("/desirePath", middleware.Pagination(desirePath.Handler))

	// Authentication
	login := &handlers.Login{}
	server.e.GET("/login", login.GetHandler)
	server.e.POST("/login", login.PostHandler, middleware.Authenticate(userRepository))

	server.e.GET("/home", func(c echo.Context) error {
		return c.String(http.StatusOK, "You are authenticated")
	}, middleware.JWTValidator)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start server
	go func() {
		if err := server.e.Start("localhost"+server.config.Server.Port); err != nil && err != http.ErrServerClosed {
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
