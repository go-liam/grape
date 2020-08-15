package page

import (
	"github.com/go-liam/util/response"
	models "grape/internal/pkg/model"
	"log"
	"testing"
)

func TestSrvPage_Create(t *testing.T) {
	item := new(Model)
	item.LanguageID = 1
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Title = "title"
	//item.ID = uuid.AutoInt64ID()
	item.Content = "Content"
	item.Description = "Description"
	item.Author = "Author"
	item.MenuID = 1
	item.SiteID = 1
	v, err := new(SrvPage).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvPage_FindOne(t *testing.T) {
	v, err := new(SrvPage).FindOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvPage_FindMulti(t *testing.T) {
	page := &response.Pagination{PageSize: 10, Current: 1}
	s := &response.ListParameter{WhereSt: " and 1=1 ", OrderSt: " order by id "}
	v, err := new(SrvPage).FindMulti(page, s)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvPage_Update(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx-page-up"}`
	item.Title = "title-page-up"
	item.ID = 2
	v, err := new(SrvPage).Update(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvPage_UpdateState(t *testing.T) {
	item := new(Model)
	item.ID = 2
	item.Status = 1
	v, err := new(SrvPage).UpdateStatus(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
