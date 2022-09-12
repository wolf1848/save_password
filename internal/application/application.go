package application

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"startGo/internal/handler"
	"startGo/internal/repository"
	"startGo/internal/service"
)

func initConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
func test(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Run() error {

	confErr := initConfig()
	if confErr != nil {
		return confErr
	}

	repository, dbErr := repository.CreateRepository()
	if dbErr != nil {
		return dbErr
	}

	service := service.CreateService(repository)

	handler := handler.CreateHandler(service)

	serverErr := handler.Run(os.Getenv("SERVER_PORT"))
	if serverErr != nil {
		return serverErr
	}

	return nil
}
