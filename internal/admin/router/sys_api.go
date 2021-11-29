package router

import (
	"github.com/cyj19/go-web/internal/admin/api/v1/api"
	"github.com/cyj19/go-web/internal/admin/global"
	"github.com/cyj19/go-web/internal/admin/store"
	"github.com/cyj19/go-web/internal/pkg/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// InitApiRouter 注册接口路由
func InitApiRouter(r *gin.RouterGroup, factoryIns store.Factory, authMiddleware *jwt.GinJWTMiddleware) {
	apiv1 := r.Group("/api")
	apiv1.Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(factoryIns, global.Conf, global.Enforcer))
	router2 := r.Group("/api").Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(factoryIns, global.Conf, global.Enforcer),
		middleware.Idempotence(global.RedisIns, global.Conf.Server.IdempotenceTokenName))
	{
		apiHandler := api.NewSysApiHandler(factoryIns)
		// 创建操作要增加幂等性校验
		router2.POST("/add", apiHandler.Create)
		apiv1.DELETE("/delete", apiHandler.BatchDelete)
		apiv1.PATCH("/update", apiHandler.Update)
		apiv1.POST("/page", apiHandler.GetPage)
		apiv1.POST("/list", apiHandler.GetList)

	}
}
