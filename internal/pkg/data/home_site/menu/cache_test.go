package menu

import (
	"log"
	"testing"
)

func TestSrvMenu_CacheMulti(t *testing.T) {
	v, err := new(SrvMenu).CacheMulti()
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
