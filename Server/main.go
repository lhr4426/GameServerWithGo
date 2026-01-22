package main

import (
	"game-server/internal/logger"
	"game-server/network"
)

func main() {
	logger.Init()
	logger.NetworkLogger.Println("Server Start")

	network.StartServer()
}
