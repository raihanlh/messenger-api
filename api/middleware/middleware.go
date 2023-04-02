package middleware

import (
	"github.com/labstack/echo/v4"
)

type middlewares struct {
	// Add your dependency here, example:
	// usecase *dependency.Usecases
}

type Middlewares interface {
	AuthToken(next echo.HandlerFunc) echo.HandlerFunc
	Authenticate(next echo.HandlerFunc) echo.HandlerFunc
}

func New(e *echo.Echo) Middlewares {
	return &middlewares{}
}
