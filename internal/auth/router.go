package auth

import (
	"go-web/internal/auth/api/v1/user"
	"go-web/internal/auth/store/mysql"

	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	installAPI(g)
}

func installAPI(g *gin.Engine) {
	factoryIns, err := mysql.GetMySQLFactory()
	if err != nil {
		panic(err)
	}

	//初始化数据库表
	err = mysql.MigrateTable()
	if err != nil {
		panic(err)
	}

	v1 := g.Group("/v1")
	{
		userHandler := user.NewUserHandler(factoryIns)
		v1.POST("/auth", userHandler.Token)
	}
}
