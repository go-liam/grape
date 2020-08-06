package svlog

import (
	"github.com/go-liam/util/conv"
	"testing"
)

func TestInit(t *testing.T) {
	Init("config")
}

func TestWriteMsg(t *testing.T) {
	WriteLog(conv.TimeNowInt64(), "msg", LevelInformational)
}

//func TestDestroy(t *testing.T) {
//	Destroy()
//}
//
//func TestFlush(t *testing.T) {
//	Flush()
//}
