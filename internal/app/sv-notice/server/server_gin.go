package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/conv"
	"github.com/go-liam/util/response"
	"grape/internal/pkg/middleware/metric"
	"net/http"
	"os"
)

const (
	//partGin    = ":7405"
	serverName = "sv-notice"
	path       = "/" + serverName
)

func RunServerGin(port string) {
	println(serverName+":run api partGin ", port)
	engine := gin.New()
	// 设置路由
	SetupRouter(engine)
	engine.Run(port)
}

func SetupRouter(engine *gin.Engine) {
	g := engine.Group(path)
	engine.GET("/", Index)
	engine.GET(path, Index)
	// 监控
	if metric.Server.IsOpen {
		g.Use(metric.Server.NewMetric())
		engine.GET(path+"/metrics/v2", gin.WrapH(metric.Server.Handler()))
	} else {
		engine.GET(path+"/metrics/v2", metric.Server.HandlerClose)
	}
	//设置路由中间件
	g.GET("/v1/:name", configCore)
}

func Index(c *gin.Context) {
	hostName, _ := os.Hostname()
	c.String(http.StatusOK, serverName+",5-4-2136 running on "+hostName)
}

func configCore(c *gin.Context) {
	name := c.Param("name")
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
	var obj map[string]interface{}
	conv.JsonStringToStruct(jsonSt, &obj)
	c.JSON(http.StatusOK, response.APIResponseData(0, "OK", obj))
}
