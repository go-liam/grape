package test

import (
	"grape/configs/testdata"
	"grape/internal/app/api_cms/config"
	"grape/internal/app/api_cms/router"
	router3 "grape/internal/pkg/middleware/router"
	"testing"
)

func TestAPI_Index(t *testing.T) {
	back, _ := router3.RouteTestTool(router.SetupRouterTest(), "GET", "/", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestAPI_v1_Index(t *testing.T) {
	back, _ := router3.RouteTestTool(router.SetupRouterTest(), "GET", config.Path+"/v1", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestAPI_v1_demos(t *testing.T) {
	back, _ := router3.RouteTestTool(router.SetupRouterTest(), "GET", config.Path+"/api/v1/demos", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}
