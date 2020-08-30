package test

import (
	"grape/configs/testdata"
	"grape/internal/app/api-cms/config"
	"grape/internal/app/api-cms/router"
	router3 "grape/internal/pkg/middleware/router"
	"strings"
	"testing"
)

func TestAPI_v1_Login(t *testing.T) {
	data := `
{
  "name": "name",
"client":1,
"time":60,
  "password": "123456"
}
	`
	reader := strings.NewReader(data)
	back, _ := router3.RouteTestTool(router.SetupRouterTest(), "POST", config.Path+"/api/v1/login", reader, testdata.UserKey, testdata.AdminKey)
	println(back)
}

func TestAPI_v1_Permissions(t *testing.T) {
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/cms/v1/user/permissions", nil, testdata.UserKey)
	println(back)
}
