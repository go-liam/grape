package role

import (
	"log"
	"testing"
)

func TestSrvRole_Create(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Name = "title-menu-2"
	//item.ID = uuid.AutoInt64ID()
	v, err := new(SrvRole).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvRole_FindOne(t *testing.T) {
	v, err := new(SrvRole).FindOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvRole_FindMulti(t *testing.T) {
	v, err := new(SrvRole).FindMultiByNil()
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvRole_Update(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Name = "title"
	item.ID = 1
	v, err := new(SrvRole).Update(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvRole_UpdateState(t *testing.T) {
	item := new(Model)
	item.ID = 1
	item.Status = 1
	v, err := new(SrvRole).UpdateStatus(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestName_FindMultiByIDs(t *testing.T) {
	v1, v2 := Server.FindMultiByIDs([]int64{1, 2})
	log.Printf("v:=%+v\n", v1)
	log.Printf("v2:=%+v\n", v2)
}
