package recharge

import (
	"github.com/go-liam/util/response"
	"log"
	"testing"
)

func TestSrvRecharge_Create(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Remark = "title"
	//item.ID = uuid.AutoInt64ID()
	v, err := new(SrvRecharge).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvRecharge_FindOne(t *testing.T) {
	v, err := new(SrvRecharge).FindOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvRecharge_FindMulti(t *testing.T) {
	page := &response.Pagination{PageSize: 10, Current: 1}
	s := &response.ListParameter{WhereSt: " and 1=1 ", OrderSt: " order by id "}
	v, err := new(SrvRecharge).FindMulti(page, s)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("page:=%+v\n", page)
}

func TestSrvRecharge_Update(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx2"}`
	item.Remark = "title"
	item.ID = 1
	v, err := new(SrvRecharge).Update(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvRecharge_UpdateState(t *testing.T) {
	item := new(Model)
	item.ID = 1
	item.Status = 1
	v, err := new(SrvRecharge).UpdateStatus(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
