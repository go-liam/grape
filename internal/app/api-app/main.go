package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/tracing"
	"github.com/go-liam/tracing/config"
	"grape/pkg/middleware/limiter"
	"grape/pkg/middleware/metric"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var (
	port       = ":7201"
	serverName = "api-app"
)

func main() {
	println("run api port ", port)
	engine := gin.New()
	// 设置路由
	SetupRouter(engine)
	engine.Run(port)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}

func SetupRouter(engine *gin.Engine) {
	path := "/" + serverName + ""
	//链路跟踪
	tracing.Init(config.DefaultConfig)
	engine.Use(tracing.NewGinMiddlewareHandle(serverName))
	g := engine.Group(path)
	// 限流
	limit := limiter.NewLimiter(limiter.MaxNumber)
	engine.Use(limiter.LimitHandler(limit))
	// 监控
	if metric.Server.IsOpen {
		g.Use(metric.Server.NewMetric())
		engine.GET(path+"/metrics/v2", gin.WrapH(metric.Server.Handler()))
	} else {
		engine.GET(path+"/metrics/v2", metric.Server.HandlerClose)
	}

	engine.GET("/", Index)
	engine.GET(path, Index)
	g.GET("/v1/:name", Index)
	//404
	engine.NoRoute(func(c *gin.Context) {
		c.String(404, "请求方法不存在 ")
	})
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello,It works.index "+serverName)
}
