package user

import (
	"github.com/go-liam/util/response"
	models "grape/internal/pkg/model"
	"log"
	"testing"
)

func TestSrvUser_Create(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Remark = "title"
	//item.ID = uuid.AutoInt64ID()
	v, err := new(SrvUser).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvUser_FindOne(t *testing.T) {
	v, err := new(SrvUser).FindOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvUser_FindMulti(t *testing.T) {
	page := &response.Pagination{PageSize: 10, Current: 1}
	s := &models.ListParameter{WhereSt: " and 1=1 ", OrderSt: " order by id "}
	v, err := new(SrvUser).FindMulti(page, s)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("page:=%+v\n", page)
}

func TestSrvUser_Update(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx2"}`
	item.Remark = "title"
	item.ID = 1
	item.Phone = "136000000"
	v, err := new(SrvUser).Update(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvUser_UpdateState(t *testing.T) {
	item := new(Model)
	item.ID = 1
	item.Status = 1
	v, err := new(SrvUser).UpdateStatus(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
