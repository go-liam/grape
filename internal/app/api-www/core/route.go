package core

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/tracing"
	"grape/internal/pkg/middleware/metric"
	"grape/internal/pkg/middleware/trace"
	"net/http"
)

func SetupRouter(engine *gin.Engine) {
	path := "/"
	g := engine.Group(path)
	g.GET("/", Index)
	g.GET("ip", ClientIPHeader)
	//404
	engine.NoRoute(func(c *gin.Context) {
		c.String(404, "请求方法不存在 ")
	})
}

func SetupRouterMetric(engine *gin.Engine) {
	g := engine.Group(Path)
	//engine.Use(router.AccessControlAllow())
	//链路跟踪
	tracing.Init(trace.DefaultConfig)
	engine.Use(tracing.NewGinMiddlewareHandle(ServerName))
	// 限流
	//limit := limiter.NewLimiter(limiter.MaxNumber)
	//engine.Use(limiter.LimitHandler(limit))
	// 监控
	if metric.Server.IsOpen {
		g.Use(metric.Server.NewMetric())
		engine.GET(Path+"/metrics/v2", gin.WrapH(metric.Server.Handler()))
	} else {
		engine.GET(Path+"/metrics/v2", metric.Server.HandlerClose)
	}
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "1-Hello,It works.index "+ServerName)
}
