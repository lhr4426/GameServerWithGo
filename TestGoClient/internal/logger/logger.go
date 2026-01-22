package logger

import (
	"log"
	"os"
)

var ClientLogger *log.Logger

func Init() {

	// log.New(출력, 접두사, 맨 앞에 포맷 제어)
	ClientLogger = log.New(
		os.Stdout,
		"[Network] ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile,
	)
}
