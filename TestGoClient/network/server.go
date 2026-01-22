package network

import (
	"game-client/internal/logger"
	"net"
)

// import (
// 	"game-server/dispatch"
// )

// 1. 연결 기다림
// 2. 연결 들어오면 루프 만들어서 Connection 개체 만들기

func StartServer() {
	// 1. 9000번 포트로 리스닝 시작
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		logger.ClientLogger.Fatalln("Server Listening Error : ", err.Error())
	}

	logger.ClientLogger.Println("Server Listening on Port 9000")

	for {
		rawConn, err := ln.Accept() // 2. 클라이언트가 들어올 때 까지 블로킹
		if err != nil {
			logger.ClientLogger.Println("Accept Error: ", err)
			continue
		}

		logger.ClientLogger.Println("Client Connected : ", rawConn.RemoteAddr().String())

		conn := NewConnection(rawConn)

		go conn.ReadLoop()
	}
}
