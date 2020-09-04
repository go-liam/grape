package test

import (
	"grape/configs/testdata"
	"grape/internal/app/api_cms/config"
	"grape/internal/app/api_cms/router"
	router3 "grape/internal/pkg/middleware/router"
	"strings"
	"testing"
)

func TestCMS_v1_power_list(t *testing.T) {
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/cms/v1/rbac/powers", nil, testdata.UserKey)
	println(back)
}

func TestCMS_v1_role_info(t *testing.T) {
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/cms/v1/rbac/role/1", nil, testdata.UserKey)
	println(back)
}

func TestCMS_v1_role_list(t *testing.T) {
	p := "?pageSize=20&current=1"
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/cms/v1/rbac/roles"+p, nil, testdata.UserKey)
	println(back)
}

func TestCMS_v1_role_update(t *testing.T) {
	data := `
{
  "extended": {"info":"xx-info"},
  "name":"name--休闲鞋"
}
	`
	reader := strings.NewReader(data)
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "PUT", config.Path+"/cms/v1/rbac/role/3", reader, testdata.UserKey)
	println(back)
}

func TestCMS_v1_role_delete(t *testing.T) {
	data := `
{
  "ids":["0","3","4","5"]
}
	`
	reader := strings.NewReader(data)
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "PUT", config.Path+"/cms/v1/rbac/role-status", reader, testdata.UserKey)
	println(back)
}

func TestCMS_v1_role_updatePower(t *testing.T) {
	data := `
{
  "power_ids":"1,2,3"
}
	`
	reader := strings.NewReader(data)
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "PUT", config.Path+"/cms/v1/rbac/role-power/3", reader, testdata.UserKey)
	println(back)
}

func TestCMS_v1_role_Add(t *testing.T) {
	data := `
{
  "extended": {"info":"xx-info"},
  "name":"name--休闲鞋",  
	"power_ids": [1,2,3]
}
	`
	reader := strings.NewReader(data)
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "POST", config.Path+"/cms/v1/rbac/role", reader, testdata.UserKey)
	println(back)
}