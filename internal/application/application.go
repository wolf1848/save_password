package application

import (
	"os"
	"startGo/internal/handler"
	"startGo/internal/repository"
	"startGo/internal/service"

	"github.com/joho/godotenv"
)

func initConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
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

	serverErr := handler.Run(":" + os.Getenv("SERVER_PORT"))
	if serverErr != nil {
		return serverErr
	}

	return nil
}
