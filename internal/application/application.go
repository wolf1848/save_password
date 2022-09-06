package application

import (
	"github.com/joho/godotenv"
	"os"
	"startGo/internal/repository"
	"startGo/internal/server"
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

	db, dbErr := repository.PoolCreate()
	if dbErr != nil {
		return dbErr
	}

	s := server.InitServer(db)

	s.Server.GET("/testSql", s.GetUser)

	serverErr := s.Run(os.Getenv("SERVER_PORT"))
	if serverErr != nil {
		return serverErr
	}

	return nil
}
