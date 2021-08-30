package router

import (
	"go-web/internal/admin/api/v1/user"
	"go-web/internal/admin/global"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 注册用户路由
func InitUserRouter(r *gin.RouterGroup, factoryIns store.Factory, authMiddleware *jwt.GinJWTMiddleware) {

	userv1 := r.Group("/user")
	userv1.Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(factoryIns, global.Conf, global.Enforcer))
	{
		userHandler := user.NewSysUserHandler(factoryIns)
		userv1.GET("/info", userHandler.GetUserInfo)
		userv1.POST("/add", userHandler.Create)
		userv1.DELETE("/delete", userHandler.BatchDelete)
		userv1.PATCH("/update", userHandler.Update)
		userv1.PATCH("/role/update", userHandler.UpdateRoleForUser)
		userv1.POST("/page", userHandler.GetPage)

	}

}
