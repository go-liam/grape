package config

import (
	"log"
	"testing"
)

func TestSrvSite_CacheOne(t *testing.T) {
	v, err := new(SrvConfig).CacheOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
