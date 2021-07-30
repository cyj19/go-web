package router

import (
	"go-web/internal/admin/api/v1/user"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// 注册用户路由
func InitUserRouter(r *gin.RouterGroup, factoryIns store.Factory) {

	// 除了公共路由，其他路由都需要权限验证
	userv1 := r.Group("/user")
	userv1.Use(middleware.CasbinMiddleware(initialize.GetEnforcerIns()))
	{
		userHandler := user.NewUserHandler(factoryIns)
		userv1.GET("/:name", userHandler.GetByUsername)

		userv1.POST("/list", userHandler.List)

		userv1.POST("/add", userHandler.Create)

		userv1.POST("/page", userHandler.GetPage)

		userv1.POST("/role", userHandler.SetUserRole)
	}

}
