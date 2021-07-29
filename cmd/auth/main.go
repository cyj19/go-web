package main

import (
	"context"
	"fmt"
	"go-web/internal/auth"
	"go-web/internal/pkg/initialize"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 初始化配置
	initialize.Config("auth.dev.yml", "auth.prod.yml")

	// 初始化日志

	// 初始化MySQL
	initialize.MySQL()

	// 初始化Redis
	initialize.Redis()

	// 初始化路由
	g := auth.Router()
	configuration := initialize.GetConfiguration()
	host := "0.0.0.0"
	port := configuration.Server.Port

	// 服务器启动
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: g,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// 打印日志
			log.Fatalf("Could not listen on %s:%d: %v\n", host, port, err)
		}
	}()

	// 优雅关闭服务
	//参考地址：https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-without-context/server.go

	// 用于监听系统的打断信号
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
