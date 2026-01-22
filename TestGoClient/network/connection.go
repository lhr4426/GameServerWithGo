package network

import (
	"game-client/codec"
	"game-client/dispatch"
	"game-client/internal/logger"
	"game-client/packet"
	"game-client/protocol"
	"net"

	"google.golang.org/protobuf/proto"
)

type Connection struct {
	Conn   net.Conn
	Writer *packet.PacketWriter
	Reader *packet.PacketReader
}

// 새 Connection 객체 생성
func NewConnection(c net.Conn) *Connection {
	return &Connection{
		Conn:   c,
		Writer: packet.NewPacketWriter(),
		Reader: packet.NewPacketReader(),
	}
}

// 계속 패킷 받는 함수
func (conn *Connection) ReadLoop() {
	defer conn.Conn.Close() // 문제 생기면 알아서 연결 닫음

	buf := make([]byte, 4096)

	for {
		n, err := conn.Conn.Read(buf) // 바이트 배열에다가 읽어옴 (블로킹). 반환값은 읽어온 바이트 개수
		if err != nil {
			logger.ClientLogger.Println("Read Loop Error : ", err.Error())
			return
		}

		msg, err := codec.DeserializePacket(buf[:n])

		dispatch.Dispatch(conn, msg)
	}
}

// 패킷을 받아서 연결 정보를 통해 보냄
func (conn *Connection) SendPacket(pkt *protocol.Packet) error {
	frame, err := conn.Writer.Write(pkt)
	if err != nil {
		logger.ClientLogger.Println("Send Packet Error : ", err.Error())
		return err
	}

	_, err = conn.Conn.Write(frame)
	return err
}

// 페이로드 부분만 받아서 알아서 패킷화
func (conn *Connection) SendMessage(msgType protocol.MessageType, msg proto.Message) error {
	payload, err := proto.Marshal(msg)
	if err != nil {
		logger.ClientLogger.Println("Send Message Error : ", err.Error())
		return err
	}

	pkt := &protocol.Packet{
		Type:    msgType,
		Payload: payload,
	}

	return conn.SendPacket(pkt)
}

func (conn *Connection) Close() error {
	return conn.Conn.Close()
}
