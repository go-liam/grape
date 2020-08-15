package change_log

import (
	"github.com/go-liam/util/response"
	"log"
	"testing"
)

func TestSrvChange_Create(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Remark = "title"
	//item.ID = uuid.AutoInt64ID()
	v, err := new(SrvChange).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvChange_FindOne(t *testing.T) {
	v, err := new(SrvChange).FindOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvChange_FindMulti(t *testing.T) {
	page := &response.Pagination{PageSize: 10, Current: 1}
	s := &response.ListParameter{WhereSt: " and 1=1 ", OrderSt: " order by id "}
	v, err := new(SrvChange).FindMulti(page, s)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("page:=%+v\n", page)
}

func TestSrvChange_Update(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx2"}`
	item.Remark = "title"
	item.ID = 1
	v, err := new(SrvChange).Update(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvChange_UpdateState(t *testing.T) {
	item := new(Model)
	item.ID = 1
	item.Status = 1
	v, err := new(SrvChange).UpdateStatus(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
