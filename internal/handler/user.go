package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) CreateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Struct!")
}
