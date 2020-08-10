package redis

import (
	"log"
	"testing"
)

func TestMock_Redis_data(t *testing.T) {
	MockRedis()
	sv := Server
	n := "pm|heZi_u1"
	f := sv.IsExist(n)
	log.Println("mock-IsExist=", f)
	sv.Set(n, "mock-xxxx-xxx", 100)
	v, _ := sv.Get(n)
	log.Println("mock-v=", v)
	f2 := sv.IsExist(n)
	log.Println("mock-IsExist=", f2)
	sv.Delete(n)
	f3 := sv.IsExist(n)
	log.Println("mock-IsExist=", f3)
}
