package middleware

import (
	"bluebell/render"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

func TokenBucketWithWaitMiddleware(r rate.Limit, b int, timeout time.Duration) gin.HandlerFunc {
	limiters := &sync.Map{}
	return func(c *gin.Context) {
		// 获取限速器
		key := c.ClientIP() // key 除了 ip 之外也可以是其他的，例如 header，user name 等
		l, _ := limiters.LoadOrStore(key, rate.NewLimiter(r, b))

		// 这里注意不要直接使用 gin 的 context 默认是没有超时时间的
		ctx, cancel := context.WithTimeout(c, timeout)
		defer cancel()

		if err := l.(*rate.Limiter).Wait(ctx); err != nil {
			// 这里先不处理日志了，如果返回错误就直接 429
			render.ResponseAbort(c, http.StatusTooManyRequests, "Too Many Requests")
			return
		}
		c.Next()
	}
}

func TokenBucketWithAllowMiddleware(r rate.Limit, b int) gin.HandlerFunc {
	limiters := &sync.Map{}
	return func(c *gin.Context) {
		// 获取限速器
		key := c.ClientIP() // key 除了 ip 之外也可以是其他的，例如 header，user name 等
		l, _ := limiters.LoadOrStore(key, rate.NewLimiter(r, b))

		if !l.(*rate.Limiter).Allow() {
			// 这里先不处理日志了，如果返回错误就直接 429
			render.ResponseAbort(c, http.StatusTooManyRequests, "Too Many Requests")
			return
		}
		c.Next()
	}
}

func LeakBucketMiddleware(rps int) gin.HandlerFunc {
	limiters := &sync.Map{}
	return func(c *gin.Context) {
		// 获取限速器
		key := c.ClientIP() // key 除了 ip 之外也可以是其他的，例如 header，user name 等
		l, _ := limiters.LoadOrStore(key, ratelimit.New(rps))
		_ = l.(ratelimit.Limiter).Take()
		c.Next()
	}
}
