package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"grape/pkg/middleware/metric"
	"grape/pkg/util"
	"grape/pkg/util/conv"
	"net/http"
	"os"
)

const (
	partGin    = ":7408"
	serverName = "sv-money"
	path = "/"+serverName
)

func RunServerGin() {
	println(serverName+":run api partGin ", partGin)
	engine := gin.New()
	// 设置路由
	SetupRouter(engine)
	engine.Run(partGin)
}

func SetupRouter(engine *gin.Engine) {
	g := engine.Group(path)
	// 监控
	if metric.Server.IsOpen {
		g.Use(metric.Server.NewMetric())
		engine.GET(path+"/metrics/v2", gin.WrapH(metric.Server.Handler()))
	} else {
		engine.GET(path+"/metrics/v2", metric.Server.HandlerClose)
	}
	//设置路由中间件
	engine.GET("/",Index )
	engine.GET(path,Index )
	g.GET("v1/:name",configCore )
}

func Index(c *gin.Context) {
	hostName, _ := os.Hostname()
	c.String(http.StatusOK, serverName+", running on " +hostName)
}

func configCore(c *gin.Context) {
	name:= c.Param("name")
	jsonSt := `
{
  "name": "%s",
  "mysql": [
    {
      "id": 123456789012,
      "host": "localhost:3307",
      "name": "box2api",
      "user": "root",
      "psw": "123456"
    }
  ]
}
`
	jsonSt = fmt.Sprintf(jsonSt, name)
	var obj  map[string]interface{}
	conv.JsonStringToStruct(jsonSt, &obj)
	c.JSON(http.StatusOK, util.APIResponseData(0,"OK",obj))
}