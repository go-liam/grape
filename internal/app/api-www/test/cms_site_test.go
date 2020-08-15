package test

import (
	"grape/configs/testdata"
	router2 "grape/internal/app/api-www/router"
	"grape/internal/pkg/middleware/router"
	"strings"
	"testing"
)

var (
	//urlSite = "/www/v1/"
	urlCms = "/www/cms/v1/"
)

func TestCMSWWW_Site_Info(t *testing.T) {
	back, _ := router.RouteTestTool(router2.SetupRouterTest(), "GET", urlCms+"site/1", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestCMSWWW_Site_list(t *testing.T) {
	back, _ := router.RouteTestTool(router2.SetupRouterTest(), "GET", urlCms+"sites", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestCMSWWW_Site_Create(t *testing.T) {
	data := `
{
  "name": "laborum-name",
  "description": "description exercitation commodo nisi ea",
  "teacherID": 3,
  "channelID": 2,
  "auditorNum" :10,
  "totalNum":20
}
	`
	reader := strings.NewReader(data)
	back, _ := router.RouteTestTool(router2.SetupRouterTest(), "POST", urlCms+"site", reader, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestCMSWWW_Site_Update(t *testing.T) {
	back, _ := router.RouteTestTool(router2.SetupRouterTest(), "PUT", urlCms+"site", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestCMSWWW_Site_updateStatus(t *testing.T) {
	back, _ := router.RouteTestTool(router2.SetupRouterTest(), "PATCH", urlCms+"site", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}
