package main

import (
	"encoding/binary"
	"fmt"
	"game-server/protocol"

	"google.golang.org/protobuf/proto"
)

func main() {
	// Frame 1
	frame1, _ := SerializePacket(makePingPacket(1))
	// Frame 2
	frame2, _ := SerializePacket(makePingPacket(2))
	// Frame 3 (일부만)
	frame3, _ := SerializePacket(makePingPacket(3))

	// 일부러 쪼갠다
	combined := append(frame1, frame2...)
	combined = append(combined, frame3[:5]...) // 미완성

	frames, remain := ParseFrames(combined)

	for i, frame := range frames {
		fmt.Println("\n\n---- Frame ", i, " ----")

		fmt.Println(frame)

		pkt, err := DeserializePacket(frame)
		if err != nil {
			fmt.Println("Deserialize Error : ", err)
			continue
		}

		fmt.Printf("\nPacket Type : %v", pkt.Type)

		switch pkt.Type {
		case protocol.MessageType_PING:
			var ping protocol.Ping
			if err := proto.Unmarshal(pkt.Payload, &ping); err != nil {
				fmt.Println("Ping Unmarshal Error : ", err)
			}
			fmt.Printf(" | Ping Timestamp : %v", ping.Timestamp)
		}
	}

	fmt.Println("\n\nRemain Frames : ", remain)
}

func makePingPacket(ts int64) *protocol.Packet {
	ping := &protocol.Ping{Timestamp: ts}
	payload, _ := proto.Marshal(ping)

	return &protocol.Packet{
		Type:    protocol.MessageType_PING,
		Payload: payload,
	}
}

// Buffer에 쌓인 byte 배열을 제대로 수신된 프레임들과 남은 프레임으로 쪼개서 반환
func ParseFrames(buffer []byte) ([][]byte, []byte) {
	var frames [][]byte
	offset := 0

	for {
		// 1. 앞에 붙는 길이 헤더를 검사함. 4보다 작으면 지금 처리 X
		if len(buffer[offset:]) < 4 {
			break
		}

		length := binary.BigEndian.Uint32(buffer[offset : offset+4])
		frameSize := int(4 + length)

		// 2. 프레임이 다 안왔으면 지금 처리 X
		if len(buffer[offset:]) < frameSize {
			break
		}

		// 3. 프레임 추출해서 프레임들에다가 넣음
		frame := buffer[offset : offset+frameSize]
		frames = append(frames, frame)

		offset += frameSize
	}

	remain := buffer[offset:] // 처리 못한것들
	return frames, remain
}

// 패킷 구조체를 바이트 배열로 변경
func SerializePacket(pkt *protocol.Packet) ([]byte, error) {
	// 1. 패킷을 바이트 배열로 변경
	packetBytes, err := proto.Marshal(pkt)
	if err != nil {
		return nil, err
	}

	// 2. 빈 바이트 배열을 만듬
	// 이때 빈 바이트 배열의 길이는 4 + 바이트 배열로 변경한 패킷의 길이
	frame := make([]byte, 4+len(packetBytes))

	// 3. 빈 바이트 배열의 앞 4바이트를 바이트 배열로 변경한 패킷 길이로 정함
	binary.BigEndian.PutUint32(frame[:4], uint32(len(packetBytes)))

	// 4. 뒤에 바이트 배열로 변경한 패킷을 붙임
	copy(frame[4:], packetBytes)

	return frame, nil
}

// 역직렬화 가능한 byte 배열을 받고, 패킷으로 반환
func DeserializePacket(frame []byte) (*protocol.Packet, error) {
	// 여기까지 왔으면 일단 역직렬화 가능한 바이트 배열이 들어왔다고 가정
	// 역직렬화 해서 Packet 구조체로 변환 후, 문제 없으면 반환함
	var pkt protocol.Packet
	if err := proto.Unmarshal(frame[4:], &pkt); err != nil {
		return nil, err
	}

	return &pkt, nil
}
