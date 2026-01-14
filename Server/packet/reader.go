package packet

import "encoding/binary"

type PacketReader struct {
	buffer []byte // 참고 : 변수명이 대문자로 시작하면 public, 아니면 private임
}

// PacketReader 초기화. 빈 버퍼를 포함함
func NewPacketReader() *PacketReader {
	return &PacketReader{buffer: make([]byte, 0)}
}

// PacketReader에 남은 버퍼에 새 데이터를 붙여서 패킷화 시도
func (r *PacketReader) Pull(data []byte) ([][]byte, error) {
	frames, remain := ParseFrames(append(r.buffer, data...))
	r.buffer = remain
	return frames, nil
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
