package handlers

import (
	"grape/configs/testdata"
	"grape/internal/app/api-www/core"
	"grape/internal/pkg/middleware/router"
	"testing"
)

func TestAPIWWWW_Index(t *testing.T) {
	back, _ := router.RouteTestTool(core.SetupRouterTest(), "GET", "/", nil, testdata.UserKey, testdata.AdminKey)
	println(back)
}
