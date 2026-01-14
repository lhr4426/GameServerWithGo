package main

import (
	"game-server/protocol"

	"google.golang.org/protobuf/proto"
)

func main() {

}

func makePingPacket(ts int64) *protocol.Packet {
	ping := &protocol.Ping{Timestamp: ts}
	payload, _ := proto.Marshal(ping)

	return &protocol.Packet{
		Type:    protocol.MessageType_PING,
		Payload: payload,
	}
}
