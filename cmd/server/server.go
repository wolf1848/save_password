package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.Static("/", "web/front/dist")

	if err := e.Start(":3125"); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
