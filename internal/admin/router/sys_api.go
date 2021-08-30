package router

import (
	"go-web/internal/admin/api/v1/api"
	"go-web/internal/admin/global"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 注册接口路由
func InitApiRouter(r *gin.RouterGroup, factoryIns store.Factory, authMiddleware *jwt.GinJWTMiddleware) {
	apiv1 := r.Group("/api")
	apiv1.Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(factoryIns, global.Conf, global.Enforcer))
	{
		apiHandler := api.NewSysApiHandler(factoryIns)

		apiv1.POST("/add", apiHandler.Create)
		apiv1.DELETE("/delete", apiHandler.BatchDelete)
		apiv1.PATCH("/update", apiHandler.Update)
		apiv1.POST("/page", apiHandler.GetPage)
		apiv1.POST("/list", apiHandler.GetList)

	}
}
