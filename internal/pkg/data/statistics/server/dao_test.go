package server

import (
	"github.com/go-liam/util/response"
	"log"
	"testing"
)

func TestSrvServer_Create(t *testing.T) {
	item := new(Model)
	item.Description = "test-Description-a"
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Title = "title"
	item.URL = "http://localhost:4001/api_app/api/v1/health/site"
	//item.ID = 2 //uuid.AutoInt64ID()
	v, err := new(SrvServer).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvServer_FindOne(t *testing.T) {
	v, err := new(SrvServer).FindOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvServer_FindMulti(t *testing.T) {
	page := &response.Pagination{PageSize: 10, Current: 1}
	s := &response.ListParameter{WhereSt: " and 1=1 ", OrderSt: " order by id "}
	v, err := new(SrvServer).FindMulti(page, s)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvServer_Update(t *testing.T) {
	item := new(Model)
	item.Description = "test-Description-au-server"
	item.Extended = `{"x":1,"b":"xxx3"}`
	item.Title = "title-server"
	item.URL = "http://localhost:4001/api_app/api/v1/health/site"
	item.ID = 1
	v, err := new(SrvServer).Update(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvServer_UpdateState(t *testing.T) {
	item := new(Model)
	item.ID = 1
	item.Status = 1
	v, err := new(SrvServer).UpdateStatus(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvServer_UpdateHealth(t *testing.T) {
	item := new(Model)
	item.RunStatus = 1
	item.Health = `{"x":1,"b":"xxx3"}`
	item.ID = 1
	v, err := new(SrvServer).UpdateHealth(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
