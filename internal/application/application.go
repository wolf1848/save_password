package application

import "startGo/internal/server"

func Run() {
	s := new(server.Server)
	s.Run(":8891")
}
