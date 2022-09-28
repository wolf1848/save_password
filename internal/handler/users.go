package handler

import (
	"fmt"
	"net/http"
	"startGo/internal/data"

	"github.com/labstack/echo/v4"
)

type Res struct {
	Message string
	Error   string
	Data    []*data.User
}

func (h *Handler) UserGroupRoute() {
	g := h.server.Group("/users")
	g.PUT("", h.createUser)
}

func (h *Handler) createUser(c echo.Context) error {

	u := new(data.User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	res := &Res{}

	if data, err := h.service.UserService.Create(u); err != nil {
		fmt.Println(err)
		res.Message = "Возникла проблема."
		res.Error = fmt.Sprintf("Ошибка : %v", err)
		return c.JSON(http.StatusOK, res)
	} else {
		res.Message = "Пользователь успешно добавлен."
		res.Data = append(res.Data, data)
		return c.JSON(http.StatusOK, res)
	}

}
