package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api_app/config"
	"grape/internal/app/api_cms/handlers/user"
	"net/http"
)

func setupRouterAPI(engine *gin.Engine) {
	path := config.Path + "/api/v1"
	engine.GET("/", Index)
	engine.GET(path, Index)
	//other
	g := engine.Group(path)
	g.POST("/login", user.LoginGin) //登录
	//
	setupRouterAPIShop(engine)
}

func setupRouterUser(engine *gin.Engine) {
	path := config.Path + "/cms/v1"
	engine.GET(path, Index)
	//user
	//setupRouterCMSUser(engine)
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "1-Hello,It works.index "+config.ServerName)
}
