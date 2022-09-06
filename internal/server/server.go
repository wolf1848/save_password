package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"startGo/internal/route"
)

type Server struct {
	server *echo.Echo
}

func (s *Server) Run(port string) error {
	e := echo.New()

	e.GET("/users/:id", route.GetUser)

	if err := e.Start(port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
