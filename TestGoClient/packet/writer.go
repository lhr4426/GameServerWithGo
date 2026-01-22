package packet

import (
	"encoding/binary"
	"game-client/internal/logger"
	"game-client/protocol"

	"google.golang.org/protobuf/proto"
)

type PacketWriter struct{}

// PacketWriter 초기화
func NewPacketWriter() *PacketWriter {
	return &PacketWriter{}
}

// 직렬화 된 패킷 반환
func (w *PacketWriter) Write(pkt *protocol.Packet) ([]byte, error) {
	return SerializePacket(pkt)
}

// 패킷 구조체를 바이트 배열로 변경
func SerializePacket(pkt *protocol.Packet) ([]byte, error) {
	// 1. 내부에서 돌던 패킷을 Protobuf에 규정한 패킷으로 바꿈
	protoPkt := &protocol.Packet{
		Type:    pkt.Type,
		Payload: pkt.Payload,
	}

	// 2. 패킷을 바이트 배열로 변경
	packetBytes, err := proto.Marshal(protoPkt)
	if err != nil {
		logger.ClientLogger.Println("Serialize Packet Error : ", err.Error())
		return nil, err
	}

	// 3. 빈 바이트 배열을 만듬
	// 이때 빈 바이트 배열의 길이는 4 + 바이트 배열로 변경한 패킷의 길이
	frame := make([]byte, 4+len(packetBytes))

	// 4. 빈 바이트 배열의 앞 4바이트를 바이트 배열로 변경한 패킷 길이로 정함
	binary.BigEndian.PutUint32(frame[:4], uint32(len(packetBytes)))

	// 5. 뒤에 바이트 배열로 변경한 패킷을 붙임
	copy(frame[4:], packetBytes)

	return frame, nil
}
