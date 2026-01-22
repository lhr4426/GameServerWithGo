package dispatch

import (
	"game-client/common"
	"reflect"
)

// 핸들러는 무조건 두 정보를 가져야 함 (연결 정보 + 메시지)
type HandlerFunc func(conn common.ConnContext, msg any)

// MessageType에 대응하는 핸들러를 연결하는 딕셔너리(맵)
var handlers = map[reflect.Type]HandlerFunc{}

// reflect.Type은 런타임에 확인하는 타입임

// 핸들러 딕셔너리에 등록
func Register(msg any, handler HandlerFunc) {
	t := reflect.TypeOf(msg)
	handlers[t] = handler
}

// 핸들러 딕셔너리를 기반으로 디스패치
func Dispatch(conn common.ConnContext, msg any) {
	t := reflect.TypeOf(msg)
	handler := handlers[t]
	handler(conn, msg)
}
