package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"net/http"
	"sync"
)

func main() {
	e := gin.Default()
	e.Use(NewLimiter(3)) // 新建一个限速器，允许突发 3 个并发
	e.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	e.Run(":8085")
}

func NewLimiter(rps int) gin.HandlerFunc {
	limiters := &sync.Map{}
	return func(c *gin.Context) {
		// 获取限速器
		key := c.ClientIP() // key 除了 ip 之外也可以是其他的，例如 header，user name 等
		l, _ := limiters.LoadOrStore(key, ratelimit.New(rps))
		now := l.(ratelimit.Limiter).Take()
		fmt.Printf("now: %s\n", now)
		c.Next()
	}
}
