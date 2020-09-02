package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api_cms/config"
	"grape/internal/app/api_cms/handlers/rbac/power_cms"
	"grape/internal/app/api_cms/handlers/rbac/role_cms"
	"grape/internal/pkg/middleware/router"
)

func setupRouterCMSRBAC(engine *gin.Engine) {
	path := config.Path + "/cms/v1/rbac"
	g := engine.Group(path)
	g.Use(router.AuthMiddleWareCheckToken())
	//path
	g.GET("/powers", power_cms.GetListGin)
	g.GET("/roles", role_cms.GetListGin)
	g.PUT("/role/:id", role_cms.EditGin)
	g.PUT("/role-power/:id", role_cms.EditPowerGin)

}
