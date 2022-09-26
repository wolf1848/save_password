package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id   uint32 `json:"id"`
	test string
}

func (h *Handler) UserGroupRoute() {
	g := h.server.Group("/users")
	g.POST("", h.CreateUser)
}

func (h *Handler) CreateUser(c echo.Context) error {

	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	fmt.Println(u.Id)
	fmt.Println(u.test)

	return c.String(http.StatusOK, "Hello, Struct!")
}
