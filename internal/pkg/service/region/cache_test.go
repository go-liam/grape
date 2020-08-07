package region

import (
	"log"
	"testing"
)

func TestSrvRegion_CacheOne(t *testing.T) {
	v, err := new(SrvRegion).CacheOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
