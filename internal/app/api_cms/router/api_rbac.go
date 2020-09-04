package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api_cms/config"
	"grape/internal/app/api_cms/handlers/rbac/power_cms"
	"grape/internal/app/api_cms/handlers/rbac/role_cms"
	"grape/internal/app/api_cms/handlers/rbac/user_cms"
	"grape/internal/pkg/middleware/router"
)

func setupRouterCMSRBAC(engine *gin.Engine) {
	path := config.Path + "/cms/v1/rbac"
	g := engine.Group(path)
	g.Use(router.AuthMiddleWareCheckToken())
	//power
	g.GET("/powers", power_cms.GetListGin)
	g.POST("/power", power_cms.AddGin)
	g.PUT("/power/:id", power_cms.EditGin)
	g.PUT("/power-status", power_cms.UpdateStatusGin)
	// role
	g.GET("/role/:id", role_cms.GetInfoGin)
	g.GET("/roles", role_cms.GetListGin)
	g.PUT("/role/:id", role_cms.EditGin)
	g.PUT("/role-power/:id", role_cms.EditPowerGin)
	g.PUT("/role-status", role_cms.UpdateStatusGin)
	g.POST("/role", role_cms.AddGin)
	//user
	g.POST("/user", user_cms.AddGin)
}
