package site

import (
	"github.com/go-liam/util/conv"
	"grape/internal/pkg/data/statistics/server"
	"log"
	"testing"
)

func TestName_GetHttp(t *testing.T) {
	url := "http://localhost:4001/api_app/api/v1/health/site"
	db, err := GetHttp(url)
	log.Printf("db=%+v\n", conv.StructToJsonString(db))
	log.Printf("err=%+v\n", err)
}

func TestName_GetAllSiteHealth(t *testing.T) {
	GetAllSiteHealth()
}

func TestName_getOneSiteHealth(t *testing.T) {
	item := &server.Model{ID: 1, URL: "http://localhost:4001/api_app/api/v1/health/site"}
	getOneSiteHealth(item)
}
