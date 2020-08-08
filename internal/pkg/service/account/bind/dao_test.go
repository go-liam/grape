package bind

import (
	"grape/configs/testdata"
	"log"
	"testing"
)

func TestSrvBind_Create(t *testing.T) {
	item := new(ModelBind)
	//item.ID = uuid.AutoInt64ID()
	item.UserID = 1
	item.Name = testdata.ConstWantString
	item.Type = 1
	v, err := new(SrvBind).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvBind_FindOne(t *testing.T) {
	v, err := new(SrvBind).FindOne(1, testdata.ConstWantString)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
