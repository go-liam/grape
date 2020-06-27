package redis

import (
	"log"
	"testing"
)

func TestPKG_Redis_data(t *testing.T) {
	sv := Server
	n := "parent-1|heZi_u1"
	f := sv.IsExist(n)
	log.Println("IsExist=", f)
	sv.Set(n, "xxxx-xxx", 100)
	v, _ := sv.Get(n)
	log.Println("v=", v)
	f2 := sv.IsExist(n)
	log.Println("IsExist=", f2)
	sv.Delete(n)
	f3 := sv.IsExist(n)
	log.Println("IsExist=", f3)
}