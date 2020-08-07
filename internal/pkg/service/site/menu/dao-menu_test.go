package menu

import (
	"github.com/go-liam/util/uuid"
	"log"
	"testing"
)

func TestSrvMenu_Create(t *testing.T) {
	item := new(Model)
	item.LanguageID = 1
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Name = "title"
	item.ID = uuid.AutoInt64ID()
	v, err := new(SrvMenu).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvMenu_FindOne(t *testing.T) {
	v, err := new(SrvMenu).FindOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvMenu_FindMulti(t *testing.T) {
	v, err := new(SrvMenu).FindMulti()
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvMenu_Update(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Name = "title"
	item.ID = 1
	v, err := new(SrvMenu).Update(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvMenu_UpdateState(t *testing.T) {
	item := new(Model)
	item.ID = 1
	item.Status = 1
	v, err := new(SrvMenu).UpdateStatus(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
