package test

import (
	"grape/configs/testdata"
	"grape/internal/app/api-cms/config"
	"grape/internal/app/api-cms/router"
	router3 "grape/internal/pkg/middleware/router"
	"strings"
	"testing"
)

func TestCMS_v1_Login(t *testing.T) {
	data := `
{
  "username": "name",
"client":1,
"login_flag":0,
  "password": "123456"
}
	`
	reader := strings.NewReader(data)
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "POST", config.Path+"/api/v1/login", reader, "")
	println(back)
}

func TestCMS_v1_Permissions(t *testing.T) {
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/cms/v1/user/permissions", nil, testdata.UserKey)
	println(back)
}

func TestCMS_v1_update(t *testing.T) {
	data := `
{
 "id": "1234567890123456",
  "nickname": "nickname",
"email":"email",
  "avatar": "avatar"
}
	`
	reader := strings.NewReader(data)
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "PUT", config.Path+"/cms/v1/user", reader, testdata.UserKey)
	println(back)
}

func TestCMS_v1_Register(t *testing.T) {
	data := `
{
 "id": "1234567890123456",
  "group_ids": "1,2,3",
"remark":"remark",
  "admin": 1
}
	`
	reader := strings.NewReader(data)
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "POST", config.Path+"/cms/v1/user/register", reader, testdata.UserKey)
	println(back)
}

func TestCMS_v1_ChangePassword(t *testing.T) {
	data := `
{
 "id": "1234567890123456",
  "old_password": "old_password",
"new_password":"new_password"
}
	`
	reader := strings.NewReader(data)
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "PUT", config.Path+"/cms/v1/user/change_password", reader, testdata.UserKey)
	println(back)
}

func TestCMS_v1_RefreshGin(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEyMzQ1Njc4OTAxMjM0NTYsInV0eSI6MSwibGciOjAsImV4cCI6MTYwMTUzNjA1OX0.T2NIr1rZkwrqPE80ahMNl5YewBCUnDcN7LekuwwW9MY"
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/cms/v1/user/refresh", nil, token)
	println(back)
}

func TestCMS_v1_InformationGin(t *testing.T) {
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/cms/v1/user/information", nil, testdata.JWTKey)
	println(back)
}
