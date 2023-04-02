package middleware

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/raihanlh/messenger-api/internal/app/dependency"
)

type middlewares struct {
	// Add your dependency here, example:
	usecases *dependency.Usecases
}

type Middlewares interface {
	AuthToken(next echo.HandlerFunc) echo.HandlerFunc
	Authenticate(next echo.HandlerFunc) echo.HandlerFunc
}

func New(e *echo.Echo, u *dependency.Usecases) Middlewares {
	return &middlewares{
		usecases: u,
	}
}
