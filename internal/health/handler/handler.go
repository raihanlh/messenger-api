package handler

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/raihanlh/messenger-api/internal/health"
)

type handler struct{}

func New() health.Handler {
	return &handler{}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce plain
// @Success 200 {object} string
// @Router /health [get]
func (h *handler) CheckHealth(e echo.Context) error {
	return e.String(200, "OK")
}
