package configclient

import (
	"log"
	"testing"
)

func TestInfo(t *testing.T) {
	v1, v2 := Info("sv-ai", 0)
	log.Printf("v1=%+v\n", v1)
	log.Printf("v1=%+v\n", v2)
}

func TestList(t *testing.T) {
	v1, v2 := List("sv-ai")
	log.Printf("v1=%+v\n", v1)
	log.Printf("v1=%+v\n", v2)
	for _, v := range v1 {
		log.Printf("v =%+v\n", v)
	}
}
