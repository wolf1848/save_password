package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) UserGroupRoute() {
	g := h.server.Group("/user")
	g.PUT("", h.CreateUser)
}

func (h *Handler) CreateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Struct!")
}
