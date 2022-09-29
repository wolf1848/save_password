package handler

import (
	"net/http"
	"os"
	"startGo/internal/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
	server  *echo.Echo
}

func (h *Handler) initRoute() {

	//Описать  middleware для чека авторизации
	//h.server.Use()

	if debug, err := strconv.ParseBool(os.Getenv("SERVER_DEBUG")); err == nil {
		h.server.Debug = debug
	}
	h.UserGroupRoute()
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
