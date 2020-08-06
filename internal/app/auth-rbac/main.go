package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/tracing"
	"github.com/go-liam/tracing/config"
	"grape/internal/app/auth-rbac/core"
	"grape/pkg/middleware/limiter"
	"grape/pkg/middleware/metric"
	"log"
	"os"
	"os/signal"
)

const (
	port       = ":7302"
	serverName = "auth-rbac"
)

func main() {
	path := "/" + serverName
	engine := gin.Default()
	//链路跟踪
	tracing.Init(config.DefaultConfig)
	engine.Use(tracing.NewGinMiddlewareHandle(serverName))
	engine.GET("/", core.Index)
	engine.GET(path, core.Index)
	g := engine.Group(path + "/v1/")
	// 限流
	limit := limiter.NewLimiter(limiter.MaxNumber)
	engine.Use(limiter.LimitHandler(limit))
	//g := engine.Group(path+"/v1/")
	g.GET("/", core.Index)
	g.GET("user/:userID", core.UserInfo)
	g.GET("check/:userID", core.CheckUserInfo)
	// 监控
	if metric.Server.IsOpen {
		g.Use(metric.Server.NewMetric())
		engine.GET(path+"/metrics/v2", gin.WrapH(metric.Server.Handler()))
	} else {
		engine.GET(path+"/metrics/v2", metric.Server.HandlerClose)
	}
	engine.Run(port)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}
