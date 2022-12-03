package logger

import (
	"log"
	"os"
	"time"
)

var Log *log.Logger
var Debug bool

func init() {
	var file *os.File
	var err error
	if Debug {
		file = os.Stdout
	} else {
		if file, err = os.OpenFile("CurrencyBotLog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755); err != nil {
			file.Close()
			panic(err)
		}
	}
	Log = log.New(file, time.Now().String(), 0)
}

type checker interface {
	Check()
}
type A struct {
}

func (a A) Check() {

}
