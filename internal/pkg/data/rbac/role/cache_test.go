package role

import (
	"log"
	"testing"
)

func TestSrvRole_CacheMulti(t *testing.T) {
	v, err := new(SrvRole).CacheMulti()
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
