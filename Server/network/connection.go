package network

import (
	"game-server/packet"
	"game-server/protocol"
	"net"

	"google.golang.org/protobuf/proto"
)

type Connection struct {
	Conn   net.Conn
	Writer *packet.PacketWriter
	Reader *packet.PacketReader
}

// 패킷을 받아서 연결 정보를 통해 보냄
func (conn *Connection) SendPacket(pkt *protocol.Packet) error {
	frame, err := conn.Writer.Write(pkt)
	if err != nil {
		return err
	}

	_, err = conn.Conn.Write(frame)
	return err
}

// 페이로드 부분만 받아서 알아서 패킷화
func (conn *Connection) SendMessage(msgType protocol.MessageType, msg proto.Message) error {
	payload, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	pkt := &protocol.Packet{
		Type:    msgType,
		Payload: payload,
	}

	return conn.SendPacket(pkt)
}
