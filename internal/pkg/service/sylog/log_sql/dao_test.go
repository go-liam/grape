package log_sql

import (
	"github.com/go-liam/util/response"
	"github.com/go-liam/util/uuid"
	models "grape/internal/pkg/model"
	"log"
	"testing"
)

func TestSrvLog_Create(t *testing.T) {
	item := new(Model)
	item.Extended = `{"x":1,"b":"xxx"}`
	item.Msg = "title"
	item.ID = uuid.AutoInt64ID()
	item.Level = 0
	v, err := new(SrvLog).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvLog_FindMulti(t *testing.T) {
	page := &response.Pagination{PageSize: 10, Current: 1}
	s := &models.ListParameter{WhereSt: " and 1=1 ", OrderSt: " order by id "}
	v, err := new(SrvLog).FindMulti(page, s)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("page:=%+v\n", page)
}

func TestSrvLog_Delete(t *testing.T) {
	ti := int64(1694367853)
	v, err := new(SrvLog).Delete(ti)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
