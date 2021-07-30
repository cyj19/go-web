package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"go-web/internal/admin/router"
	"go-web/internal/admin/store/mysql"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/middleware"
)

// 初始化路由
func Router() *gin.Engine {
	// 创建一个没有中间件的路由
	g := gin.New()

	// 初始化go-jwt中间件
	authMiddleware, err := middleware.InitGinJWTMiddleware(nil)
	if err != nil {
		panic(fmt.Sprintf("初始化jwt中间件失败：%v", err))
	}

	factoryIns, err := mysql.GetMySQLFactory()
	if err != nil {
		panic(fmt.Sprintf("初始化工厂实例失败：%v", err))
	}

	configuration := initialize.GetConfiguration()
	apiRouter := g.Group(configuration.Server.UrlPrefix)
	v1 := apiRouter.Group(configuration.Server.ApiVersion)
	// 所有路由都需要token校验
	v1.Use(authMiddleware.MiddlewareFunc())
	router.InitPublicRouter(v1)           // 注册公共路由
	router.InitUserRouter(v1, factoryIns) // 注册用户路由
	router.InitRoleRouter(v1, factoryIns) // 注册角色路由
	router.InitMenuRouter(v1, factoryIns) // 注册菜单路由
	router.InitApiRouter(v1)              //注册接口路由
	return g
}
