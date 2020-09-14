package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api_cms/config"
	"grape/internal/app/api_cms/handlers/demo"
	"grape/internal/app/api_cms/handlers/health"
	"grape/internal/app/api_cms/handlers/user"
	"net/http"
)

func setupRouterAPI(engine *gin.Engine) {
	path := config.Path + "/api/v1"
	engine.GET("/", health.Index)
	engine.GET(path, health.Index)
	g := engine.Group(path)
	// health
	g.GET("/health/site", health.Site)
	//other
	g.POST("/login", user.LoginGin) //登录
	g.GET("/demos", demo.GetListGin)

}

func setupRouterCMS(engine *gin.Engine) {
	path := config.Path + "/cms/v1"
	engine.GET(path, Index)
	//user
	setupRouterCMSUser(engine)
	//rbac
	setupRouterCMSRBAC(engine)
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "1-"+config.ServerName)
}
