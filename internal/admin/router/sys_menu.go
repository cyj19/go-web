package router

import (
	"github.com/cyj19/go-web/internal/admin/api/v1/menu"
	"github.com/cyj19/go-web/internal/admin/global"
	"github.com/cyj19/go-web/internal/admin/store"
	"github.com/cyj19/go-web/internal/pkg/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// InitMenuRouter 注册菜单路由
func InitMenuRouter(r *gin.RouterGroup, factoryIns store.Factory, authMiddleware *jwt.GinJWTMiddleware) {

	menuv1 := r.Group("/menu")
	menuv1.Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(factoryIns, global.Conf, global.Enforcer))
	router2 := r.Group("/menu").Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(factoryIns, global.Conf, global.Enforcer),
		middleware.Idempotence(global.RedisIns, global.Conf.Server.IdempotenceTokenName))
	{
		menuHandler := menu.NewSysMenuHandler(factoryIns)
		// 创建操作要增加幂等性校验
		router2.POST("/add", menuHandler.Create)
		menuv1.DELETE("/delete", menuHandler.BatchDelete)
		menuv1.PATCH("/update", menuHandler.Update)
		menuv1.POST("/list", menuHandler.GetList)
		menuv1.GET("/all", menuHandler.GetMenusByRoleId)
		menuv1.POST("/page", menuHandler.GetPage)
	}
}
