package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/tracing"
	"github.com/go-liam/tracing/config"
	"grape/internal/app/auth-jwt/single"
	"grape/internal/app/auth-jwt/sv"
	"grape/pkg/middleware/limiter"
	"grape/pkg/middleware/metric"
	"log"
	"net/http"
	"os"
	"os/signal"
)

const (
	//part         = ":7301"
	tokenName    = "token"
	tokenTime    = 3600 * 24 * 365
	cookieDomain = "mygs.com"
)

var (
	port       = ":7301"
	serverName = "auth-jwt"
	path =  "/"+serverName
)

var badUid = []int64{1,2,3,119} // 黑名单


func main() {
	engine := gin.Default()
	engine.GET("/", Index)
	engine.GET(path,Index )
	RouterUser(engine)
	RouterServer(engine)
	engine.Run(port)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    "index page auth-jwt",
	})
}

func RouterUser(engine *gin.Engine)  {
	g1 := engine.Group(path+"/v1/")
	//链路跟踪
	tracing.Init(config.DefaultConfig)
	g1.Use(tracing.NewGinMiddlewareHandle(serverName))
	// 限流
	limit := limiter.NewLimiter(limiter.MaxNumber)
	g1.Use(limiter.LimitHandler(limit))
	// 监控
	if metric.Server.IsOpen {
		g1.Use(metric.Server.NewMetric())
		engine.GET(path+"/metrics/v2", gin.WrapH(metric.Server.Handler()))
	} else {
		engine.GET(path+"/metrics/v2", metric.Server.HandlerClose)
	}
	g1.GET("login", single.Login)
	g1.GET("info", single.CheckUser)
	g1.GET("refresh", single.Refresh)
	g1.GET("logout", single.Logout)
	g1.GET("check", single.CheckBadUser)
}

func RouterServer(engine *gin.Engine)  {
	g := engine.Group(path+"/sv/v1/")
	g.GET("login", sv.Login)
	g.GET("info", sv.CheckUser)
	g.GET("refresh", sv.Refresh)
	g.GET("logout", sv.Logout)
	g.GET("check", sv.CheckBadUser)
	//404
	engine.NoRoute(func(c *gin.Context) {
		c.String(404, "请求方法不存在 " )
	})
}