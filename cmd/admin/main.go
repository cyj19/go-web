package main

import (
	"context"
	"fmt"
	"go-web/internal/admin"
	"go-web/internal/admin/global"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/model"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	// 处理初始化阶段可能抛出的panic，目的并非要恢复程序，而是在程序退出前记录错误信息和堆栈信息
	defer func() {
		if err := recover(); err != nil {
			// 如果初始化日志完成，则写入日志文件
			if global.Log != nil {
				// 把错误信息和堆栈信息写入日志文件
				global.Log.Fatal(ctx, "未知异常，退出程序", err, string(debug.Stack()))
			} else {
				log.Fatalf("未知异常，退出程序: %v   %s", err, string(debug.Stack()))
			}

		}
	}()

	// 初始化配置文件
	global.Box, global.Conf = initialize.Config("admin.dev.yml", "admin.prod.yml")

	// 初始化日志
	global.Log = initialize.InitLogger(global.Conf)
	global.Log.Info(ctx, "初始化日志完成...")

	// 初始化MySQL
	global.DB = initialize.MySQL(global.Conf.Mysql, global.Log, new(model.SysUser), new(model.SysRole), new(model.SysMenu), new(model.SysCasbin), new(model.SysApi))
	global.Log.Info(ctx, "初始化mysql完成...")

	// 初始化Redis
	global.RedisIns = initialize.Redis(global.Conf.Redis)
	global.Log.Info(ctx, "初始化redis完成...")

	// 初始化Casbin
	global.Enforcer = initialize.Casbin(global.DB, global.Box, global.Conf)
	global.Log.Info(ctx, "初始化casbin完成...")

	// 初始化操作工厂
	factoryIns, err := store.NewSqlFactory(global.DB)
	if err != nil {
		panic(fmt.Sprintf("初始化工厂实例失败：%v", err))
	}

	// 初始化数据
	admin.InitData(ctx, factoryIns)
	global.Log.Info(ctx, "初始化数据完成...")

	// 初始化路由
	g := admin.Router(ctx, factoryIns)

	host := "0.0.0.0"
	port := global.Conf.Server.Port
	//启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: g,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Fatal(ctx, "Could not listen on %s:%d: %v\n", host, port, err)
		}
	}()

	// 优雅关闭服务
	//参考地址：https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-without-context/server.go

	//监听系统打断信号
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	global.Log.Info(ctx, "Shutting down server...")

	// 留5秒用于处理未完成的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Log.Error(ctx, "Server forced to shutdown: %v", err)
	}

}
