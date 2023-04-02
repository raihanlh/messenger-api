package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	apiMiddleware "gitlab.com/raihanlh/messenger-api/api/middleware"
	"gitlab.com/raihanlh/messenger-api/api/router"
	"gitlab.com/raihanlh/messenger-api/config"
	"gitlab.com/raihanlh/messenger-api/internal/app"
	"gitlab.com/raihanlh/messenger-api/pkg/logger"
	"gitlab.com/raihanlh/messenger-api/pkg/validator"
)

// @title Messenger API
// @version 1.0
// @description Messenger API server.
// @termsOfService http://swagger.io/terms/

// @contact.name Raihan Luthfi
// @contact.url https://github.com/raihanlh
// @contact.email raihan.luthfi.h@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath
func main() {
	conf := config.New()
	logger.Setup(conf)

	databases := app.NewDatabases(conf)
	repositories := app.NewRepositories(databases)
	usecases := app.NewUsecases(repositories)
	handlers := app.NewHandlers(usecases)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.Validator = validator.New()

	mw := apiMiddleware.New(e)
	router.MapRoutes(e, handlers, mw)

	log.Fatal(e.Start(":" + conf.Port))
}
