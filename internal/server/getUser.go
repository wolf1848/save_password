package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func (s *Server) GetUser(c echo.Context) error {

	var greeting string
	err := s.db.Pool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
	// User ID from path `users/:id`
	//id := c.Param("id")
	return c.String(http.StatusOK, greeting)
}
