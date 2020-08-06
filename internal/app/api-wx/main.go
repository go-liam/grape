package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-liam/tracing"
	"github.com/go-liam/tracing/config"
	"github.com/go-liam/util/request"
	"grape/internal/pkg/middleware/limiter"
	"grape/internal/pkg/middleware/metric"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var (
	port       = ":7204"
	serverName = "api-wx"
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
	path := "/" + serverName
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
	g.GET("/v1/:name", helloHandler)
	//404
	engine.NoRoute(func(c *gin.Context) {
		c.String(404, "请求方法不存在 ")
	})
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello,It works.index "+serverName)
}

func helloHandler(c *gin.Context) {
	log.Printf("head: %+v\n", c.Request.Header)
	name, _ := os.Hostname()
	path := c.Request.URL.Path
	st1 := "%s: host: %s, path: %s ,token: %s ,UserID: %s ,ip:%s"
	st := fmt.Sprintf(st1, serverName, name, path, c.Request.Header.Get("Token"), c.Request.Header.Get("UserID"), request.ClientIP(c.Request))
	c.String(http.StatusOK, st)
}
