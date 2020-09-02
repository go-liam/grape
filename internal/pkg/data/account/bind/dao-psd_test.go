package bind

import (
	"github.com/go-liam/util/encrypt/md5"
	"log"
	"testing"
)

func TestSrv_Create(t *testing.T) {
	item := new(ModelPassword)
	item.ID = 1
	item.Password = md5.GetMD5Hash(KeyMd5 + "123456")
	v, err := ServerPsd.Create(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrv_FindOne(t *testing.T) {
	v, err := ServerPsd.FindOne(1)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrv_Updte(t *testing.T) {
	item := new(ModelPassword)
	item.ID = 1
	item.Password = md5.GetMD5Hash(KeyMd5 + "123456")
	v, err := ServerPsd.Update(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}

func TestSrv_CheckAndCreate(t *testing.T) {
	item := new(ModelPassword)
	//item.ID = uuid.AutoInt64ID()
	item.ID = 1
	item.Password = md5.GetMD5Hash(KeyMd5 + "123456")
	v, err := ServerPsd.CheckAndCreate(item)
	log.Printf("v:=%+v\n", v)
	log.Printf("err:=%+v\n", err)
}
