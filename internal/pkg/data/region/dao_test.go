package region

import (
	"github.com/go-liam/util/response"
	"github.com/go-liam/util/uuid"
	//models "grape/internal/pkg/model"
	"log"
	"testing"
)

func TestSrvRegion_Create(t *testing.T) {
	item := new(Model)
	item.RegionID = 1
	item.UserID = uuid.AutoInt64ID()
	v, err := new(SrvRegion).Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("item:=%+v\n", item) // item.ID = 53
}

func TestSrvRegion_FindOne(t *testing.T) {
	v, err := new(SrvRegion).FindOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrvRegion_FindMulti(t *testing.T) {
	page := &response.Pagination{PageSize: 10, Current: 1}
	s := &response.ListParameter{WhereSt: " and 1=1 ", OrderSt: " order by user_id "}
	v, err := new(SrvRegion).FindMulti(page, s)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
	log.Printf("page:=%+v\n", page)
}
