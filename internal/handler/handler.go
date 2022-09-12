package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"startGo/internal/service"
)

type Handler struct {
	service *service.Service
	server  *echo.Echo
}

func (h *Handler) initRoute() {
	h.server.GET("/", h.CreateUser)
}

func CreateHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
		server:  echo.New(),
	}
}

func (h *Handler) Run(port string) error {

	h.initRoute()

	if err := h.server.Start(port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
