package main

import (
	"github.com/labstack/echo/v4"
	"startGo/internal/route"
)

func router(e *echo.Echo) {
	e.Static("/", "web/front/dist")
	e.GET("/users/:id", route.GetUser)
}
