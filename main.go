package main

import (
	"bluebell/logger"
	"bluebell/router"
	"bluebell/setting"
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title bluebell
// @version 1.0
// @description gin 框架开发 web 应用测试程序
// @termsOfService http://swagger.io/terms/
//
// @contact.name author:zorro
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
//
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host 192.168.24.133:8081
// @BasePath /api/v1/
func main() {
	cfgFile := flag.String("f", "./conf/config.yaml", "指定配置文件路径")

	flag.Parse()

	//1.加载配置
	if err := setting.Init(*cfgFile); err != nil {
		panic(fmt.Errorf("init settings failed,err:%v \n", err))
	}

	//2.初始化日志
	if err := logger.Init(setting.Cfg().Log, setting.Cfg().Mode); err != nil {
		panic(fmt.Errorf("init logger failed,err:%v \n", err))
	}
	defer zap.L().Sync()

	////3.初始化 MySQL 连接
	//if err := mysql.Init(setting.Cfg().MySQL); err != nil {
	//	panic(fmt.Errorf("init mysql failed,err:%v \n", err))
	//}
	//defer mysql.Close()
	//
	////4.初始化 Redis 连接
	//if err := redis.Init(setting.Cfg().Redis); err != nil {
	//	panic(fmt.Errorf("init redis failed,err:%v \n", err))
	//}
	//defer redis.Close()
	//
	//// 雪花算法生成分布式ID
	//if err := snowflake.Init(setting.Cfg().StartTime, setting.Cfg().MachineID); err != nil {
	//	panic(fmt.Sprintf("init snowflake failed, err:%v\n", err))
	//	return
	//}

	// 注册路由
	r := router.SetupRouter()

	//6.启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Cfg().Port),
		Handler: r,
	}

	go func() {
		// 开启一个 goroutine 启动服务
		zap.L().Info("开始监听", zap.String("port", fmt.Sprintf(":%d", setting.Cfg().Port)))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")

	// 创建一个 5 秒超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
