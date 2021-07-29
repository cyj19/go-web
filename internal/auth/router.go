package auth

import (
	"go-web/internal/auth/api/v1/user"
	"go-web/internal/auth/initialize"
	"go-web/internal/auth/store/mysql"
	"go-web/internal/pkg/global"

	"github.com/gin-gonic/gin"
)

var customServer *global.CustomServer

// 初始化路由
func initRouter() *gin.Engine {
	g := gin.Default()
	installMiddleware(g)
	installAPI(g)

	return g
}

// 安装中间件
func installMiddleware(g *gin.Engine) {

}

// 安装API
func installAPI(g *gin.Engine) {
	factoryIns, err := mysql.GetMySQLFactory()
	customServer = initialize.GetCustomServer()
	if err != nil {
		panic(err)
	}

	apiRouter := g.Group(customServer.UrlPrefix)
	v1 := apiRouter.Group(customServer.ApiVersion)
	{
		userHandler := user.NewUserHandler(factoryIns)
		v1.POST("/auth", userHandler.Token)
	}

}
