package test

import (
	"grape/configs/testdata"
	"grape/internal/app/api_app/config"
	"grape/internal/app/api_app/router"
	router3 "grape/internal/pkg/middleware/router"
	"strings"
	"testing"
)

func TestCMS_v1_Login_pwd(t *testing.T) {
	data := `
{
  "username": "root",
"client":2,
"login_flag":1,
  "password": "123456--"
}
	`
	reader := strings.NewReader(data)
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "POST", config.Path+"/api/v1/user/login-pwd", reader, "")
	println(back)
}

func TestCMS_v1_Login_sms(t *testing.T) {
	data := `
{
  "phone": "138000000000",
"client":2,
"login_flag":1,
  "code": "123456"
}
	`
	reader := strings.NewReader(data)
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "POST", config.Path+"/api/v1/user/login-sms", reader, "")
	println(back)
}

func TestCMS_v1_logout(t *testing.T) {
	token := testdata.JWTKey
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/api/v1/user/logout", nil, token)
	println(back)
}

func TestCMS_v1_RefreshGin(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEyMzQ1Njc4OTAxMjM0NTYsInV0eSI6MSwibGciOjAsImV4cCI6MTYwMTUzNjA1OX0.T2NIr1rZkwrqPE80ahMNl5YewBCUnDcN7LekuwwW9MY"
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/api/v1/user/refresh", nil, token)
	println(back)
}

func TestCMS_v1_InformationGin(t *testing.T) {
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/api/v1/user/info", nil, testdata.JWTKey)
	println(back)
}
