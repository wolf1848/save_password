package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"startGo/internal/repository"
)

type Server struct {
	Server *echo.Echo
	db     *repository.PSQL
}

func InitServer(db *repository.PSQL) *Server {
	return &Server{
		Server: echo.New(),
		db:     db,
	}
}

func (s *Server) Run(port string) error {

	if err := s.Server.Start(port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
