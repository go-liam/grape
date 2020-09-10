package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api_app/config"
	"grape/internal/app/api_app/handlers/shop"
	"grape/internal/app/api_app/handlers/sms"
)

func setupRouterAPIShop(engine *gin.Engine) {
	path := config.Path + "/api/v1/shop"
	//path
	g := engine.Group(path)
	// shop
	g.GET("/index-category", shop.CategoryGin)
	g.GET("/position", shop.PositionGin)
	g.GET("/shops", shop.ShopsGin)
	g.GET("/search-shops", shop.SearchShopsGin)
	//sms
}

func setupRouterAPISms(engine *gin.Engine) {
	path := config.Path + "/api/v1/sms"
	//path
	g := engine.Group(path)
	// shop
	g.GET("/send-code", sms.SendGin)

}
