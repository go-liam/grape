package handlers

import (
	"grape/configs/testdata"
	"grape/internal/app/api-www/core"
	"grape/internal/pkg/middleware/router"
	"testing"
)

var (
	urlSite = "/www/v1/"
)

func TestAPIWWWW_Site_Info(t *testing.T) {
	back, _ := router.RouteTestTool(core.SetupRouterTest(), "GET", urlSite+"site/1", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestAPIWWWW_Site_list(t *testing.T) {
	back, _ := router.RouteTestTool(core.SetupRouterTest(), "GET", urlSite+"sites", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestAPIWWWW_Site_Create(t *testing.T) {
	back, _ := router.RouteTestTool(core.SetupRouterTest(), "POST", urlSite+"site", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestAPIWWWW_Site_Update(t *testing.T) {
	back, _ := router.RouteTestTool(core.SetupRouterTest(), "PUT", urlSite+"site", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestAPIWWWW_Site_updateStatus(t *testing.T) {
	back, _ := router.RouteTestTool(core.SetupRouterTest(), "PATCH", urlSite+"site", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}
