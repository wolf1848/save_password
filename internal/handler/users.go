package handler

import (
	"fmt"
	"net/http"
	"startGo/internal/data"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UserGroupRoute() {
	g := h.server.Group("/users")
	g.PUT("", h.createUser)
}

func (h *Handler) createUser(c echo.Context) error {

	u := new(data.User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	_, err := h.service.UserService.Create(u)

	fmt.Println(err)

	return c.String(http.StatusOK, "Hello, Struct!")
}
