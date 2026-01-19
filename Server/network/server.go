package network

import (
	"log"
	"net"
	"time"
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
		log.Fatal(err)
	}

	log.Println(time.Now(), " | Server Listening on Port 9000")

	for {
		conn, err := ln.Accept() // 2. 클라이언트가 들어올 때 까지 블로킹
		if err != nil {
			log.Println(time.Now(), " | Accept Error: ", err)
			continue
		}

		log.Println(time.Now(), " | Client Connected : ", conn.RemoteAddr().String())
		conn.Close() // 지금 당장 처리할 건 없기 때문에 연결 종료
	}
}
