package route

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Res struct {
	Status  string `json:"status" xml:"status"`
	Message string `json:"message" xml:"message"`
}

func Default(c echo.Context) error {
	res := &Res{
		Status:  "not found",
		Message: "not element response",
	}
	return c.JSON(http.StatusOK, res)
}
