package logger

import (
	"log"
	"os"
	"time"
)

var Log *log.Logger

func init() {
	file, err := os.OpenFile("CurrencyBotLog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	Log = log.New(file, time.Now().String(), 0)
}
