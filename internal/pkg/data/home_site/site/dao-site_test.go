package site

import (
	"github.com/go-liam/util/response"
	"log"
	"testing"
)

func TestSrvSite_Create(t *testing.T) {
	item := new(Model)
	item.LanguageID = 1
	item.Description = "test-Description-a"
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Title = "title"
	item.Email = "email"
	//item.ID = 2 //uuid.AutoInt64ID()
	v, err := new(SrvSite).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvSite_FindOne(t *testing.T) {
	v, err := new(SrvSite).FindOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvSite_FindMulti(t *testing.T) {
	page := &response.Pagination{PageSize: 10, Current: 1}
	s := &response.ListParameter{WhereSt: " and 1=1 ", OrderSt: " order by id "}
	v, err := new(SrvSite).FindMulti(page, s)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvSite_Update(t *testing.T) {
	item := new(Model)
	item.Description = "test-Description-au-site"
	item.Extended = `{"x":1,"b":"xxx3"}`
	item.Title = "title-site"
	item.Email = "email-site"
	item.ID = 2
	v, err := new(SrvSite).Update(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvSite_UpdateState(t *testing.T) {
	item := new(Model)
	item.ID = 1
	item.Status = 1
	v, err := new(SrvSite).UpdateStatus(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
