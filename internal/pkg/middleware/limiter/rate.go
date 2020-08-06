package limiter

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"

	"context"
	"time"
)

// LimitHandlerB1 :等待方式，不返回错，一直等待
func LimitHandlerB1(limiter *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := limiter.Wait(context.Background()); err != nil {
			c.JSON(http.StatusOK, errInfo)
			c.Abort()
			return
		}
		c.Next()
	}
}

// LimitHandlerB2 : 没资源就返回错 ； Allow, Wait, Reserve
func LimitHandlerB2(limiter *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := limiter.Allow(); !err {
			c.JSON(http.StatusOK, errInfo)
			c.Abort()
			return
		}
		c.Next()
	}
}

func NewLimiter2(max int) *rate.Limiter {
	// 每 800ms 产生 1 个 token，最多缓存 1 个 token，如果缓存满了，新的 token 会被丢弃
	limiter := rate.NewLimiter(rate.Every(time.Duration(1000)*time.Millisecond), max)
	return limiter
}
