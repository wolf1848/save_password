package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"startGo/cmd/web/routes"
)

func main() {
	e := echo.New()

	e.GET("/users/:id", routes.GetUser)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
