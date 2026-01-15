package handler

import (
	"fmt"
	"game-server/dispatch"
	"game-server/network"
	"game-server/protocol"
	"time"
)

func init() {
	fmt.Println(time.Now(), " | Initialize Handler Started")
	dispatch.Register(&protocol.Ping{}, handlePing)
	fmt.Println(time.Now(), " | Initialize Handler Ended")
}

// MessageType : Ping 일 때 핸들러
func handlePing(conn *network.Connection, msg any) {
	_, ok := msg.(*protocol.Ping)
	// 참고 : 여기서 msg 는 any 라는 타입인데, any 는 말 그대로 뭐든 받을 수 있음
	// 타입 단언이라는 문법임. msg 안에 있는 데이터의 타입이 protocol.Ping 이면
	// 데이터를 꺼내서 req 에 넣고, 아니라면 실패함
	// 그래서 실제 데이터가 들어가는 req 와 성공 여부가 들어가는 ok 가 있음
	if !ok {
		return
	}

	fmt.Println(time.Now(), " | Ping Request")

	serverTime := time.Now().UnixMilli()
	fmt.Println(time.Now(), " | Pong Response : ", serverTime)

	pong := &protocol.Pong{
		Timestamp: serverTime,
	}

	conn.SendMessage(protocol.MessageType_PONG, pong)
}
