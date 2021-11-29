package router

import (
	"github.com/cyj19/go-web/internal/admin/api/v1/user"
	"github.com/cyj19/go-web/internal/admin/global"
	"github.com/cyj19/go-web/internal/admin/store"
	"github.com/cyj19/go-web/internal/pkg/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// InitUserRouter 注册用户路由
func InitUserRouter(r *gin.RouterGroup, factoryIns store.Factory, authMiddleware *jwt.GinJWTMiddleware) {

	userv1 := r.Group("/user")
	userv1.Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(factoryIns, global.Conf, global.Enforcer))
	router2 := r.Group("/user").Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(factoryIns, global.Conf, global.Enforcer),
		middleware.Idempotence(global.RedisIns, global.Conf.Server.IdempotenceTokenName))
	{
		userHandler := user.NewSysUserHandler(factoryIns)
		// 创建操作要增加幂等性校验
		router2.POST("/add", userHandler.Create)
		userv1.GET("/info", userHandler.GetUserInfo)
		userv1.DELETE("/delete", userHandler.BatchDelete)
		userv1.PATCH("/update", userHandler.Update)
		userv1.PATCH("/role/update", userHandler.UpdateRoleForUser)
		userv1.POST("/page", userHandler.GetPage)

	}

}
