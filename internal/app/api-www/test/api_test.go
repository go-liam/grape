package test

import (
	"grape/configs/testdata"
	router2 "grape/internal/app/api-www/router"
	"grape/internal/pkg/middleware/router"
	"testing"
)

func TestAPIWWWW_Index(t *testing.T) {
	back, _ := router.RouteTestTool(router2.SetupRouterTest(), "GET", "/", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestAPIWWWW_site_info(t *testing.T) {
	back, _ := router.RouteTestTool(router2.SetupRouterTest(), "GET", "/api_www/api/v1/site/1", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}
