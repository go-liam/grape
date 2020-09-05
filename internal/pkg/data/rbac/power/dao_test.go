package power

import (
	"github.com/go-liam/util/conv"
	"log"
	"testing"
)

func TestSrvPower_Create(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Name = "title-menu-2"
	//item.ID = uuid.AutoInt64ID()
	v, err := new(SrvPower).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvPower_FindOne(t *testing.T) {
	v, err := new(SrvPower).FindOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvPower_FindMulti(t *testing.T) {
	v, err := new(SrvPower).FindMultiByNil()
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvPower_Update(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Name = "title"
	item.ID = 1
	v, err := new(SrvPower).Update(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvPower_UpdateState(t *testing.T) {
	item := new(Model)
	item.ID = 1
	item.Status = 1
	v, err := new(SrvPower).UpdateStatus(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestName_FindMultiByIDs(t *testing.T) {
	v1, v2 := Server.FindMultiByIDs([]int64{11, 21})
	log.Printf("v:=%+v\n", conv.StructToJsonString(v1))
	log.Printf("v2:=%+v\n", v2)
}
