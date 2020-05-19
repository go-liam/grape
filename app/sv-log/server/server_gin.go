package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/util/response"
	"grape/pkg/middleware/metric"
	"grape/pkg/service/sylog"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	portGin    = ":7403"
	path = "/"+serverName
)

func RunServerGin() {
	println(serverName+":run api portGin ", portGin)
	engine := gin.New()
	// 设置路由
	SetupRouter(engine)
	engine.Run(portGin)
}

func SetupRouter(engine *gin.Engine) {
	////链路跟踪
	//tracing.Init(trace.DefaultConfig)
	//engine.Use(tracing.NewGinMiddlewareHandle(serverName))
	//设置路由中间件
	engine.GET("/", Index)
	engine.GET(path,Index )
	g := engine.Group(path)
	// 监控
	if metric.Server.IsOpen {
		g.Use(metric.Server.NewMetric())
		engine.GET(path+"/metrics/v2", gin.WrapH(metric.Server.Handler()))
	} else {
		engine.GET(path+"/metrics/v2", metric.Server.HandlerClose)
	}
	g.GET("/", Index)
	g.POST("/log", writeCore)
}

func Index(c *gin.Context) {
	hostName, _ := os.Hostname()
	c.String(http.StatusOK, serverName+", running on " +hostName)
}

type ModeWriteReq struct {
	LogID int64  `json:"logID"`
	Msg   string `json:"msg"`
	Level int32  `json:"level"`
}

func writeCore(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	req := new(ModeWriteReq)
	//println(string(body))
	json.Unmarshal(body, &req)
	if req != nil && req.LogID > 0{
		go svlog.WriteLog(req.LogID,req.Msg,req.Level)
	}
	c.JSON(http.StatusOK, response.APIResponseData(0,"OK",nil))
}