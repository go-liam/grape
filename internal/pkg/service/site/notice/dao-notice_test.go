package notice

import (
	"github.com/go-liam/util/response"
	"github.com/go-liam/util/uuid"
	models "grape/internal/pkg/model"
	"log"
	"testing"
)

func TestSrvNotice_Create(t *testing.T) {
	item := new(NoticeModel)
	item.LanguageID = 1
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Title = "title"
	item.ID = uuid.AutoInt64ID()
	item.Content = "Content"
	item.Author = "Author"
	item.SiteID = 1
	v, err := new(SrvNotice).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvNotice_FindOne(t *testing.T) {
	v, err := new(SrvNotice).FindOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvNotice_FindMulti(t *testing.T) {
	page := &response.Pagination{PageSize: 10, Current: 1}
	s := &models.ListParameter{WhereSt: " 1=1 ", OrderSt: " order by id "}
	v, err := new(SrvNotice).FindMulti(page, s)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvNotice_Update(t *testing.T) {
	item := new(NoticeModel)
	item.Extended = `{"x":1,"b":"xxx2"}`
	item.Title = "title"
	item.ID = 1
	item.Content = "ct"
	item.Author = "author"
	v, err := new(SrvNotice).Update(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvNotice_UpdateState(t *testing.T) {
	item := new(NoticeModel)
	item.ID = 1
	item.Status = 1
	v, err := new(SrvNotice).UpdateState(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
