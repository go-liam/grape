package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api_app/config"
	"net/http"
)

func setupRouterAPI(engine *gin.Engine) {
	path := config.Path + "/api/v1"
	engine.GET("/", Index)
	engine.GET(path, Index)
	//other
	setupRouterAPIShop(engine)
	setupRouterAPISms(engine)
}

func setupRouterUserCenter(engine *gin.Engine) {
	setupRouterUser(engine)
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "1-Hello,It works.index "+config.ServerName)
}
