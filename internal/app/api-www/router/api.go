package router

import (
	"github.com/gin-gonic/gin"
	"grape/internal/app/api-www/config"
	"grape/internal/app/api-www/handlers/site_cms"
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
	g.GET("/site/:id", site_cms.GetInfoGin)
	g.GET("/sites", site_cms.GetListGin)
	g.POST("/site", site_cms.AddGin)
	g.PUT("/site", site_cms.EditGin)
	g.PATCH("/site", site_cms.UpdateStatusGin)

}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "1-Hello,It works.index "+config.ServerName)
}
