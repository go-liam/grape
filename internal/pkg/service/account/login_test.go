package account

import (
	"log"
	"testing"
)

func TestName_LoginByPassword(t *testing.T) {
	v1, v2 := LoginByPassword("root", "1234568")
	log.Printf("v1=%+v\n", v1)
	log.Printf("v2=%+v\n", v2)
}
