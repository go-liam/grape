package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/tracing"
	"log"
	"grape/app/api-www/core"
	"grape/pkg/middleware/metric"
	"grape/pkg/middleware/trace"
	"net/http"
	"os"
	"os/signal"
)

const (
	port = ":7203"
	serverName ="api-www"
	path = "/"+serverName
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
	//engine.Use(router.AccessControlAllow())
	//链路跟踪
	tracing.Init(trace.DefaultConfig)
	engine.Use(tracing.NewGinMiddlewareHandle(serverName))
	engine.GET("/",Index )
	engine.GET(path,Index )
	g := engine.Group(path)
	// 限流
	//limit := limiter.NewLimiter(limiter.MaxNumber)
	//engine.Use(limiter.LimitHandler(limit))
	// 监控
	if metric.Server.IsOpen {
		g.Use(metric.Server.NewMetric())
		engine.GET(path+"/metrics/v2", gin.WrapH(metric.Server.Handler()))
	} else {
		engine.GET(path+"/metrics/v2", metric.Server.HandlerClose)
	}
	g.GET("/",Index )
	g.GET("ip",core.ClientIPHeader)
	//404
	engine.NoRoute(func(c *gin.Context) {
		c.String(404, "请求方法不存在 " )
	})
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "1-Hello,It works.index "+serverName )
}
