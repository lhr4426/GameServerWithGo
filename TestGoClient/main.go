package main

import (
	"fmt"
	"net"

	"game-client/codec"
	"game-client/packet"
	"game-client/protocol"

	"google.golang.org/protobuf/proto"
)

func main() {
	// 1. 서버 연결
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	// 2. Packet Writer / Reader
	writer := packet.NewPacketWriter()
	reader := packet.NewReader()

	// 3. Ping 생성
	ping := &protocol.Ping{}

	payload, err := proto.Marshal(ping)
	if err != nil {
		panic(err)
	}

	pkt := &protocol.Packet{
		Type:    protocol.MessageType_PING,
		Payload: payload,
	}

	// 4. Ping 전송
	if err := writer.Write(pkt); err != nil {
		panic(err)
	}

	fmt.Println("Ping sent")

	// 5. Pong 수신
	buf := make([]byte, 4096)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}

		frames, _ := reader.Feed(buf[:n])

		for _, frame := range frames {
			resp, err := codec.DeserializePacket(frame)
			if err != nil {
				continue
			}

			if resp.Type == protocol.MessageType_PONG {
				var pong protocol.Pong
				proto.Unmarshal(resp.Payload, &pong)

				fmt.Println("Pong received:", pong.Timestamp)
				return
			}
		}
	}
}
