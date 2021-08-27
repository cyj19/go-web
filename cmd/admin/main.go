package main

import (
	"context"
	"fmt"
	"go-web/internal/admin"
	"go-web/internal/admin/store/mysql"
	"go-web/internal/pkg/global"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/model"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	// 初始化配置文件
	initialize.Config(ctx, "admin.dev.yml", "admin.prod.yml")

	// 初始化日志
	initialize.InitLogger()

	// 初始化MySQL
	initialize.MySQL(new(model.SysUser), new(model.SysRole), new(model.SysMenu), new(model.SysCasbin), new(model.SysApi))

	// 初始化操作工厂
	factoryIns, err := mysql.GetMySQLFactory()
	if err != nil {
		panic(fmt.Sprintf("初始化工厂实例失败：%v", err))
	}

	// 初始化Redis
	initialize.Redis()

	// 初始化Casbin
	initialize.Casbin()
	enforcer := initialize.GetEnforcerIns()

	// 初始化数据
	admin.InitData(factoryIns, enforcer)

	// 初始化路由
	g := admin.Router(ctx, factoryIns, enforcer)

	host := "0.0.0.0"
	port := global.Conf.Server.Port
	//启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: g,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s:%d: %v\n", host, port, err)
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
