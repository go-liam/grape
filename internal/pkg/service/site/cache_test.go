package site

import (
	"log"
	"testing"
)

func TestSrvSite_CacheOne(t *testing.T) {
	v, err := new(SrvSite).CacheOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvMenu_CacheMulti(t *testing.T) {
	v, err := new(SrvMenu).CacheMulti()
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
