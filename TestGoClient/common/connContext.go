package common

import (
	"game-client/protocol"

	"google.golang.org/protobuf/proto"
)

type ConnContext interface {
	SendMessage(msgType protocol.MessageType, msg proto.Message) error
	Close() error
}
