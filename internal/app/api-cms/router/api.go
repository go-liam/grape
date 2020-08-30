package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api-cms/config"
	"grape/internal/app/api-cms/handlers/user"
	"net/http"
)

func setupRouterAPI(engine *gin.Engine) {
	path := config.Path + "/api/v1"
	engine.GET("/", Index)
	engine.GET(path, Index)
	//other
	g := engine.Group(path)
	g.POST("/login", user.LoginGin)
}

func setupRouterCMS(engine *gin.Engine) {
	path := config.Path + "/cms/v1"
	engine.GET(path, Index)
	//other
	g := engine.Group(path)
	g.GET("/", Index)
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "1-Hello,It works.index "+config.ServerName)
}
