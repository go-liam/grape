package pay

import (
	"github.com/go-liam/util/response"
	"log"
	"testing"
)

func TestSrvPay_Create(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Remark = "title"
	//item.ID = uuid.AutoInt64ID()
	v, err := new(SrvPay).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvPay_FindOne(t *testing.T) {
	v, err := new(SrvPay).FindOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvPay_FindMulti(t *testing.T) {
	page := &response.Pagination{PageSize: 10, Current: 1}
	s := &response.ListParameter{WhereSt: " and 1=1 ", OrderSt: " order by id "}
	v, err := new(SrvPay).FindMulti(page, s)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("page:=%+v\n", page)
}

func TestSrvPay_Update(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx2"}`
	item.Remark = "title"
	item.ID = 1
	v, err := new(SrvPay).Update(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvPay_UpdateState(t *testing.T) {
	item := new(Model)
	item.ID = 1
	item.Status = 1
	v, err := new(SrvPay).UpdateStatus(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
