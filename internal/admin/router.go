package admin

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"go-web/internal/admin/api/v1/user"
	"go-web/internal/admin/router"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/middleware"
)

// 初始化路由
func Router(factoryIns store.Factory, enforcer *casbin.Enforcer) *gin.Engine {
	// 创建一个没有中间件的路由
	g := gin.New()
	// 添加全局异常处理中间件
	g.Use(middleware.Exception)

	userHandler := user.NewSysUserHandler(factoryIns, enforcer)
	// 初始化go-jwt中间件
	authMiddleware, err := middleware.InitGinJWTMiddleware(userHandler.Login)
	if err != nil {
		panic(fmt.Sprintf("初始化jwt中间件失败：%v", err))
	}

	configuration := initialize.GetConfiguration()
	apiRouter := g.Group(configuration.Server.UrlPrefix)
	v1 := apiRouter.Group(configuration.Server.ApiVersion)

	//v1.Use(authMiddleware.MiddlewareFunc())
	router.InitBaseRouter(v1, authMiddleware)                       // 注册基础路由
	router.InitUserRouter(v1, factoryIns, enforcer, authMiddleware) // 注册用户路由
	router.InitRoleRouter(v1, factoryIns, enforcer, authMiddleware) // 注册角色路由
	router.InitMenuRouter(v1, factoryIns, enforcer, authMiddleware) // 注册菜单路由
	router.InitApiRouter(v1, factoryIns, enforcer, authMiddleware)  //注册接口路由
	return g
}
