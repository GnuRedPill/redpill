package utils

import (
	"log"
)

func InitLog() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
