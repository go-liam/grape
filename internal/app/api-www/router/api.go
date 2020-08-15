package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api-www/core"
	"grape/internal/app/api-www/handlers/site"
	"net/http"
)

func setupRouterAPI(engine *gin.Engine) {
	path := "/www/v1/api"
	engine.GET("/", Index)
	engine.GET(path, Index)
	//other
	g := engine.Group(path)
	g.GET("/", Index)
	//g.GET("ip", core.ClientIPHeader)
}

func setupRouterCMS(engine *gin.Engine) {
	path := "/www/cms/v1"
	engine.GET(path, Index)
	//other
	g := engine.Group(path)
	g.POST("/site", site.CreateGin)
	g.GET("/site/:id", site.GetInfoGin)
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "1-Hello,It works.index "+config.ServerName)
}
