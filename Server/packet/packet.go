package packet

import (
	"game-server/protocol"
)

type Packet struct {
	Type    protocol.MessageType
	Payload []byte
}
