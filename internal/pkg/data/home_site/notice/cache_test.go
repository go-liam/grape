package notice

import (
	"log"
	"testing"
)

func TestSrvNotice_CacheOne(t *testing.T) {
	v, err := new(SrvNotice).CacheOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
