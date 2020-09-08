package test

import (
	"grape/configs/testdata"
	"grape/internal/app/api_app/config"
	"grape/internal/app/api_app/router"
	router3 "grape/internal/pkg/middleware/router"
	"testing"
)

func TestCMS_v1_power_list(t *testing.T) {
	back, _ := router3.RouteJWTTool(router.SetupRouterTest(), "GET", config.Path+"/api/v1/shop/position", nil, testdata.UserKey)
	println(back)
}
