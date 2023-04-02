package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	httpError "gitlab.com/raihanlh/messenger-api/api/payload/http-error"
)

func (m *middlewares) AuthToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "token" {
			return next(c)
		}
		err := httpError.Forbidden("Invalid token")
		return c.JSON(http.StatusInternalServerError, err.HttpResponseError())
	}
}
