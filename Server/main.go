package main

import (
	"fmt"
	"game-server/network"
	"time"
)

func main() {
	fmt.Println(time.Now(), " | Server Start")
	network.StartServer()
}
