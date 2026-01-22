package logger

import (
	"log"
	"os"
)

var NetworkLogger *log.Logger
var DispatchLogger *log.Logger
var HandlerLogger *log.Logger
var PacketLogger *log.Logger

func Init() {

	// log.New(출력, 접두사, 맨 앞에 포맷 제어)
	NetworkLogger = log.New(
		os.Stdout,
		"[Network] ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile,
	)

	DispatchLogger = log.New(
		os.Stdout,
		"[Dispatch] ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile,
	)

	HandlerLogger = log.New(
		os.Stdout,
		"[Handler] ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile,
	)

	PacketLogger = log.New(
		os.Stdout,
		"[Packet] ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile,
	)
}
