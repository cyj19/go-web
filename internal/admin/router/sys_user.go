package router

import (
	"go-web/internal/admin/api/v1/user"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 注册用户路由
func InitUserRouter(r *gin.RouterGroup, factoryIns store.Factory, authMiddleware *jwt.GinJWTMiddleware) {

	userv1 := r.Group("/user")
	userv1.Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(initialize.GetEnforcerIns()))
	{
		userHandler := user.NewUserHandler(factoryIns)
		userv1.GET("/:name", userHandler.GetByUsername)

		userv1.POST("/list", userHandler.List)

		userv1.POST("/add", userHandler.Create)

		userv1.POST("/page", userHandler.GetPage)

		userv1.POST("/role", userHandler.SetUserRole)
	}

}
