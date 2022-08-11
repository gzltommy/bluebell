package router

import (
	"bluebell/logger"
	"bluebell/setting"
	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	sf "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

func initEngine() *gin.Engine {
	if setting.Cfg.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置成发布模式
	}

	e := gin.New()
	if setting.Cfg.PProf > 0 {
		pprof.Register(e) // 注册 pprof 相关路由
	}

	e.Use(logger.GinLogger(), logger.GinRecovery(true))
	e.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Pragma", "Cache-Control", "Connection", "Content-Length", "Content-Type", "Authorization", "X-Forwarded-For", "User-Agent", "Referer"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,

		//ExposeHeaders:          nil,
		//AllowWildcard:          false,
		//AllowBrowserExtensions: false,
		//AllowWebSockets:        false,
		//AllowFiles:             false,
		//AllowOrigins:           nil,
		//AllowOriginFunc:        nil,
	}))

	e.Use(gzip.Gzip(gzip.DefaultCompression))

	e.Use(limit.MaxAllowed(setting.Cfg.LimitConnection))

	// 最大运行上传文件大小
	e.MaxMultipartMemory = 1 << 30 // 1G

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": "404",
			"msg":  "Endpoint Not Found",
		})
	})

	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": "405",
			"msg":  "Method Not Allowed",
		})
	})

	// 注册 swagger
	//r.GET("/swagger/*any", gs.WrapHandler(sf.Handler))
	e.GET("/swagger/*any", gs.DisablingWrapHandler(sf.Handler, "NAME_OF_ENV_VARIABLE"))

	e.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return e
}
