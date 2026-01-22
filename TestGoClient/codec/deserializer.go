package codec

import (
	"game-client/internal/logger"
	"game-client/protocol"

	"google.golang.org/protobuf/proto"
)

// 역직렬화 가능한 byte 배열을 받고, 패킷으로 반환
func DeserializePacket(frame []byte) (*protocol.Packet, error) {
	// 여기까지 왔으면 일단 역직렬화 가능한 바이트 배열이 들어왔다고 가정
	// 역직렬화 해서 Protobuf Packet 구조체로 변환
	var pkt protocol.Packet
	if err := proto.Unmarshal(frame[4:], &pkt); err != nil {
		logger.ClientLogger.Println("Deserialize Packet Error : ", err)
		return nil, err
	}

	return &pkt, nil
}
