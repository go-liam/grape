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
	partGin    = ":7409"
	serverName = "sv-region"
	path       = "/" + serverName
)

func RunServerGin() {
	println(serverName+":run api partGin ", partGin)
	engine := gin.New()
	// 设置路由
	SetupRouter(engine)
	engine.Run(partGin)
}

func SetupRouter(engine *gin.Engine) {
	g := engine.Group(path + "/v1")
	// 监控
	if metric.Server.IsOpen {
		g.Use(metric.Server.NewMetric())
		engine.GET(path+"/metrics/v2", gin.WrapH(metric.Server.Handler()))
	} else {
		engine.GET(path+"/metrics/v2", metric.Server.HandlerClose)
	}
	//设置路由中间件
	engine.GET("/", Index)
	engine.GET(path, Index)
	g.GET("reg", regHandler)
	g.GET("info/:userID", userInfoHandler)
}

func Index(c *gin.Context) {
	hostName, _ := os.Hostname()
	c.String(http.StatusOK, serverName+", running on "+hostName)
}

func regHandler(c *gin.Context) {
	ip := c.Query("ip")
	rID, rInfo := IpCheck(ip)
	jsonSt := `
{
  "ip": "%s",
      "userID": "12345678901234567",
      "regionID": %d,
      "host": "%s"
}
`
	jsonSt = fmt.Sprintf(jsonSt, ip, rID, rInfo)
	var obj map[string]interface{}
	conv.JsonStringToStruct(jsonSt, &obj)
	c.JSON(http.StatusOK, util.APIResponseData(0, "OK", obj))
}

func userInfoHandler(c *gin.Context) {
	userID := c.Param("userID")
	rID, rInfo := UserIDCheck(conv.StringToInt64(userID, 0))
	jsonSt := `
{
  "userID": "%s",
      "regionID": %d,
      "host": "%s"
}
`
	jsonSt = fmt.Sprintf(jsonSt, userID, rID, rInfo)
	var obj map[string]interface{}
	conv.JsonStringToStruct(jsonSt, &obj)
	c.JSON(http.StatusOK, util.APIResponseData(0, "OK", obj))
}
