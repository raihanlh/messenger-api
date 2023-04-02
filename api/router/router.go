package router

import (
	swagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	_ "gitlab.com/raihanlh/messenger-api/api/docs"
	"gitlab.com/raihanlh/messenger-api/api/middleware"
	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
)

func MapRoutes(e *echo.Echo, h *dependency.Handlers, mw middleware.Middlewares) {
	e.GET("/swagger/*", swagger.EchoWrapHandler())
	e.GET("/health", h.Health.CheckHealth)

	api := e.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")
	user.POST("/create", h.User.Create)
	user.PATCH("/update/:id", h.User.Update, mw.AuthToken)
	user.DELETE("/delete/:id", h.User.Delete, mw.AuthToken)
	user.GET("/:id", h.User.GetById)
	user.GET("s", h.User.GetAll)

}