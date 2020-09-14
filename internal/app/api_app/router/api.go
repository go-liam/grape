package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api_app/config"
	"grape/internal/app/api_app/handlers/health"
)

func setupRouterAPI(engine *gin.Engine) {
	path := config.Path + "/api/v1"
	engine.GET("/", health.Index)
	engine.GET(path, health.Index)
	g := engine.Group(path)
	// health
	g.GET("/health/site", health.Site)
	//other
	setupRouterAPIShop(engine)
	setupRouterAPISms(engine)
}

func setupRouterUserCenter(engine *gin.Engine) {
	setupRouterUser(engine)
}

//func Index(c *gin.Context) {
//	c.String(http.StatusOK, "1-Hello,It works.index "+config.ServerName)
//}
