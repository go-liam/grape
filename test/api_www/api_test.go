package api_www

import (
	"grape/internal/app/api-www/core"
	"grape/internal/pkg/middleware/router"
	"testing"
)

func TestAPIWWWW_Index(t *testing.T) {
	back, _ := router.RouteTestTool(core.SetupRouterTest(), "GET", "/", nil, "", "")
	println(back)
}
