package health

import "github.com/labstack/echo/v4"

type Handler interface {
	CheckHealth(e echo.Context) error
}
