package admin

import (
	"context"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"go-web/internal/admin/api/v1/user"
	"go-web/internal/admin/router"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/global"
	"go-web/internal/pkg/middleware"
)

/*
	初始化路由
	参数1：上下文
	参数2：工厂实例
	参数3：casbin执行指针
*/
func Router(ctx context.Context, factoryIns store.Factory, enforcer *casbin.Enforcer) *gin.Engine {
	// 创建一个没有中间件的路由
	g := gin.New()
	// 添加访问日志
	g.Use(middleware.GinLog)
	// 添加全局异常处理中间件
	g.Use(middleware.Exception)
	// 添加全局跨域中间件
	g.Use(middleware.Cors)

	userHandler := user.NewSysUserHandler(factoryIns, enforcer)
	// 初始化go-jwt中间件
	authMiddleware, err := middleware.InitGinJWTMiddleware(userHandler.Login)
	if err != nil {
		panic(fmt.Sprintf("初始化jwt中间件失败：%v", err))
	}

	apiRouter := g.Group(global.Conf.Server.UrlPrefix)
	v1 := apiRouter.Group(global.Conf.Server.ApiVersion)

	//v1.Use(authMiddleware.MiddlewareFunc())
	router.InitBaseRouter(v1, authMiddleware)                       // 注册基础路由
	router.InitUserRouter(v1, factoryIns, enforcer, authMiddleware) // 注册用户路由
	router.InitRoleRouter(v1, factoryIns, enforcer, authMiddleware) // 注册角色路由
	router.InitMenuRouter(v1, factoryIns, enforcer, authMiddleware) // 注册菜单路由
	router.InitApiRouter(v1, factoryIns, enforcer, authMiddleware)  // 注册接口路由
	global.Log.Info(ctx, "初始化路由完成...")
	return g
}
