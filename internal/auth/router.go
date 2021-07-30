package auth

import (
	"fmt"
	"go-web/internal/auth/api/v1/user"

	"go-web/internal/auth/store/mysql"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func Router() *gin.Engine {
	g := gin.Default()

	factoryIns, err := mysql.GetMySQLFactory()
	configuration := initialize.GetConfiguration()
	if err != nil {
		panic(err)
	}
	userHandler := user.NewUserHandler(factoryIns)
	// 初始化jwt中间件
	authMiddlerware, err := middleware.InitGinJWTMiddleware(userHandler.Login)
	if err != nil {
		panic(fmt.Sprintf("初始化jwt中间件异常：%v", err))
	}
	apiRouter := g.Group(configuration.Server.UrlPrefix)
	v1 := apiRouter.Group(configuration.Server.ApiVersion)
	{
		v1.POST("/login", authMiddlerware.LoginHandler)
		v1.GET("/refresh_token", authMiddlerware.RefreshHandler)
		v1.GET("/logout", authMiddlerware.LogoutHandler)
	}

	return g
}
