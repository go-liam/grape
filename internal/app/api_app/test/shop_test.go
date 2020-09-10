package test

import (
	"grape/configs/testdata"
	"grape/internal/app/api_app/config"
	"grape/internal/app/api_app/router"
	router3 "grape/internal/pkg/middleware/router"
	"testing"
)

func TestShop_v1_position_list(t *testing.T) {
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/api/v1/shop/position", nil, testdata.UserKey)
	println(back)
}

func TestShop_v1_index_category_list(t *testing.T) {
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/api/v1/shop/index-category", nil, testdata.UserKey)
	println(back)
}

func TestShop_v1_shop_list(t *testing.T) {
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/api/v1/shop/shops", nil, testdata.UserKey)
	println(back)
}

func TestShop_v1_search_shops_list(t *testing.T) {
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/api/v1/shop/search-shops", nil, testdata.UserKey)
	println(back)
}

func TestShop_v1_sms_sendCode(t *testing.T) {
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/api/v1/sms/send-code?phone=1234567&type=1", nil, testdata.UserKey)
	println(back)
}
