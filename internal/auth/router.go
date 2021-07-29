package auth

import (
	"go-web/internal/auth/api/v1/user"

	"go-web/internal/auth/store/mysql"
	"go-web/internal/pkg/initialize"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func Router() *gin.Engine {
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
	configuration := initialize.GetConfiguration()
	if err != nil {
		panic(err)
	}

	apiRouter := g.Group(configuration.Server.UrlPrefix)
	v1 := apiRouter.Group(configuration.Server.ApiVersion)
	{
		userHandler := user.NewUserHandler(factoryIns)
		v1.POST("/auth", userHandler.Token)
	}

}
