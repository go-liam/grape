package site

import (
	"github.com/go-liam/util/conv"
	"grape/internal/pkg/data/statistics/server"
	"log"
)

func GetAllSiteHealth() {
	list, err := server.Server.CacheMulti()
	if err != nil {
		log.Printf("[ERROR]GetSiteHealth %+v\n", err)
	}
	for _, v := range list {
		getOneSiteHealth(v)
	}
}

func getOneSiteHealth(item *server.Model) {
	got, err := GetHttp(item.URL)
	if err != nil {
		log.Printf("[ERROR]GetSiteHealth %+v\n", err)
		return
	}
	item.RunStatus = got.Data.Status
	item.Health = conv.StructToJsonString(got.Data.List)
	server.Server.UpdateHealth(item)
}
