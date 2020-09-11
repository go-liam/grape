package server

import (
	"log"
	"testing"
)

func TestSrvServer_CacheOne(t *testing.T) {
	v, err := new(SrvServer).CacheOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvServer_CacheMulti(t *testing.T) {
	v, err := new(SrvServer).CacheMulti()
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
