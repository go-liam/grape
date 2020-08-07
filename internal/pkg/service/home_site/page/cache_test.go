package page

import (
	"log"
	"testing"
)

func TestSrvSite_CacheOne(t *testing.T) {
	v, err := new(SrvPage).CacheOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
