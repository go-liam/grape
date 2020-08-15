package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-liam/tracing"
	"grape/internal/app/api-www/config"
	"grape/internal/pkg/middleware/metric"
	"grape/internal/pkg/middleware/trace"
)

func SetupRouter(engine *gin.Engine) {
	setupRouterAPI(engine)
	setupRouterCMS(engine)
}

func SetupRouterMetric(engine *gin.Engine) {
	g := engine.Group(config.Path)
	//链路跟踪
	tracing.Init(trace.DefaultConfig)
	engine.Use(tracing.NewGinMiddlewareHandle(config.ServerName))
	// 限流
	//limit := limiter.NewLimiter(limiter.MaxNumber)
	//engine.Use(limiter.LimitHandler(limit))
	// 监控
	if metric.Server.IsOpen {
		g.Use(metric.Server.NewMetric())
		engine.GET(config.Path+"/metrics/v2", gin.WrapH(metric.Server.Handler()))
	} else {
		engine.GET(config.Path+"/metrics/v2", metric.Server.HandlerClose)
	}
	engine.NoRoute(func(c *gin.Context) {
		c.String(404, "请求方法不存在 ")
	})
}
