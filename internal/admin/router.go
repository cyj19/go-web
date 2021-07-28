package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"go-web/internal/admin/api/v1/menu"
	"go-web/internal/admin/api/v1/role"
	"go-web/internal/admin/api/v1/user"
	"go-web/internal/admin/common"
	"go-web/internal/admin/store/mysql"
	"go-web/internal/pkg/middleware"
)

func initRouter(g *gin.Engine) {
	initMiddleware(g)
	initAPI(g)
}

func initMiddleware(g *gin.Engine) {
	redisIns, _ := common.GetRedisIns()
	enforcer, _ := mysql.GetEnforcerIns()
	fmt.Printf("redis: %+v, enforcer: %+v \n", redisIns, enforcer)
	var notCheckTokenUrlArr, notCheckParmissionUrlArr []string
	//不需要token验证的资源
	notCheckTokenUrlArr = append(notCheckTokenUrlArr, "/v1/auth/token")
	//不需要权限验证的资源
	notCheckParmissionUrlArr = append(notCheckParmissionUrlArr, notCheckTokenUrlArr...)
	//notCheckParmissionUrlArr = append(notCheckParmissionUrlArr, "/v1/role/permission")

	g.Use(middleware.AuthMiddleware(redisIns, middleware.AllowPathPreFixSkipper(notCheckTokenUrlArr...)),
		middleware.CasbinMiddleware(enforcer, middleware.AllowPathPreFixSkipper(notCheckParmissionUrlArr...)))
}

func initAPI(g *gin.Engine) {

	factoryIns, err := mysql.GetMySQLFactory()
	if err != nil {
		panic(err)
	}
	//初始化数据库表
	mysql.MigrateTable()

	v1 := g.Group("/v1")
	{
		userHandler := user.NewUserHandler(factoryIns)
		userv1 := v1.Group("/user")
		{

			userv1.GET("/:name", userHandler.GetByUsername)

			userv1.POST("/list", userHandler.List)

			userv1.POST("/add", userHandler.Create)

			userv1.POST("/page", userHandler.GetPage)

			userv1.POST("/role", userHandler.SetUserRole)

		}
		auth := v1.Group("/auth")
		{
			auth.GET("/policy", user.LoadPolicy)
		}

		roleHandler := role.NewRoleHandler(factoryIns)
		rolev1 := v1.Group("/role")
		{
			rolev1.POST("/add", roleHandler.Create)

			rolev1.PUT("/update", roleHandler.Update)

			rolev1.DELETE("/:id", roleHandler.Delete)

			rolev1.DELETE("", roleHandler.DeleteBatch)

			rolev1.GET("/:id", roleHandler.GetById)

			rolev1.POST("/list", roleHandler.List)

			rolev1.POST("/page", roleHandler.GetPage)

			rolev1.POST("/permission", roleHandler.SetRolePermission)
		}

		menuHandler := menu.NewMenuHandler(factoryIns)
		menuv1 := v1.Group("/menu")
		{
			menuv1.POST("/add", menuHandler.Create)

			menuv1.PUT("/update", menuHandler.Update)

			menuv1.DELETE("/:id", menuHandler.Delete)

			menuv1.DELETE("", menuHandler.DeleteBatch)

			menuv1.GET("/:id", menuHandler.GetById)

			menuv1.POST("/list", menuHandler.List)

			menuv1.POST("/page", menuHandler.GetPage)
		}
	}

}
