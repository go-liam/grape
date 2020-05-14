package jwt

import (
	"github.com/go-liam/util/conv"
	"testing"
)

func TestName_badIP(t *testing.T) {
	flag := conv.ArrayInt64Contains(ServerGin.BadUidList, conv.StringToInt64("119",0))
	println("flag:",flag)
}
