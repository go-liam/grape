package power

import (
	"log"
	"testing"
)

func TestSrvSite_CacheOne(t *testing.T) {
	v, err := new(SrvPower).CacheOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
