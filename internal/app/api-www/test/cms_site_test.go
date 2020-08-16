package test

import (
	"grape/configs/testdata"
	router2 "grape/internal/app/api-www/router"
	"grape/internal/pkg/middleware/router"
	"strings"
	"testing"
)


func TestCMSWWW_Site_Info(t *testing.T) {
	back, _ := router.RouteTestTool(router2.SetupRouterTest(), "GET", urlCms+"site/1", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestCMSWWW_Site_list(t *testing.T) {
	p := "pageSize=2&current=1"
	back, _ := router.RouteTestTool(router2.SetupRouterTest(), "GET", urlCms+"sites?"+p, nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestCMSWWW_Site_Create(t *testing.T) {
	data := `
{
  "languageID": 1,
  "title": "description exercitation commodo nisi ea",
  "email": "email",
  "description": "description",
  "extended" :{"a":1,"b":"xxx"}
}
	`
	reader := strings.NewReader(data)
	back, _ := router.RouteTestTool(router2.SetupRouterTest(), "POST", urlCms+"site", reader, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestCMSWWW_Site_Update(t *testing.T) {
	data := `
{
  "languageID": 1,
  "title": "description up",
  "email": "email",
  "description": "description",
  "extended" :{"a":1,"b":"xxx"},
	"id":"1"
}
	`
	reader := strings.NewReader(data)
	back, _ := router.RouteTestTool(router2.SetupRouterTest(), "PUT", urlCms+"site", reader, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestCMSWWW_Site_updateStatus(t *testing.T) {
	data := `
{
  "IDs": ["1","2","85866711131123602"],
  "status": 1
}
	`
	reader := strings.NewReader(data)
	back, _ := router.RouteTestTool(router2.SetupRouterTest(), "PATCH", urlCms+"site", reader, testdata.UserKey, testdata.AdminKey)
	println(back)
}
