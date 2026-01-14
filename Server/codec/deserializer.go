package codec

import (
	"game-server/packet"
	"game-server/protocol"

	"google.golang.org/protobuf/proto"
)

// 역직렬화 가능한 byte 배열을 받고, 패킷으로 반환
func DeserializePacket(frame []byte) (*packet.Packet, error) {
	// 여기까지 왔으면 일단 역직렬화 가능한 바이트 배열이 들어왔다고 가정
	// 1. 역직렬화 해서 Protobuf Packet 구조체로 변환
	var pkt protocol.Packet
	if err := proto.Unmarshal(frame[4:], &pkt); err != nil {
		return nil, err
	}

	// 2. Protobuf Packet 을 우리 서버 내의 Packet 구조체로 변환
	serverPkt := &packet.Packet{
		Type:    pkt.Type,
		Payload: pkt.Payload,
	}

	return serverPkt, nil
}
