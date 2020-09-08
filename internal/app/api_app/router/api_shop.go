package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api_app/config"
	"grape/internal/app/api_app/handlers/shop"
)

func setupRouterAPIShop(engine *gin.Engine) {
	path := config.Path + "/api/v1/shop"
	//path
	g := engine.Group(path)
	// shop
	g.GET("/position", shop.PositionGin)        //查询自己信息
}
