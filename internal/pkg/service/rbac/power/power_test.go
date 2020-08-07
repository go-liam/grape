package power

import (
	"testing"
)

func TestCheckPowerByTag(t *testing.T) {
	list, _ := Server.List(1)
	tag := "box-add"
	flag := CheckPowerByTag(tag, list)
	println("flag:", flag)
}

//func TestMain(m *testing.M) {
//	log.Println("begin----power")
//	Server = ServerMock
//	m.Run()
//	log.Println("end---power")
//}

func TestGetPowerIDByTag(t *testing.T) {
	list, _ := Server.List(1)
	tag := "box-add"
	flag := GetPowerIDByTag(tag, list)
	println("flag:", flag)
}
