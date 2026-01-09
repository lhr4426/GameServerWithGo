package main

import (
	"fmt"
	"game-server/protocol"
	"time"

	"google.golang.org/protobuf/proto"
)

func main() {

	// 1. Ping 메시지 만들기

	ping := &protocol.Ping{
		Timestamp: time.Now().UnixMilli(),
	}

	// 2. Ping 을 []byte로 만들기

	payload, err := proto.Marshal(ping)
	if err != nil {
		panic(err)
	}

	// 3. []byte 로 만든 Ping을 공통 패킷으로 감싸기

	packet := &protocol.Packet{
		Type:    protocol.MessageType_PING,
		Payload: payload,
	}

	// 4. 감싼 공통 패킷을 직렬화 (실제 네트워크로 보내지는건 이게 보내짐)

	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}

	// ==== 서버가 해당 핑 메시지를 받았다고 치고 처리해보기 ====

	var recvPacket protocol.Packet
	err = proto.Unmarshal(data, &recvPacket)
	if err != nil {
		panic(err)
	}

	switch recvPacket.Type {
	case protocol.MessageType_PING:
		var recvPing protocol.Ping
		_ = proto.Unmarshal(recvPacket.Payload, &recvPing)

		fmt.Println("받은 핑 : ", recvPing.Timestamp)

	default:
		fmt.Println("모르는 메시지 타입이 들어옴.")
	}

}
