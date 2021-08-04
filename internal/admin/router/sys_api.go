package router

import (
	"go-web/internal/admin/api/v1/api"
	"go-web/internal/admin/store"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 注册接口路由
func InitApiRouter(r *gin.RouterGroup, factoryIns store.Factory, authMiddleware *jwt.GinJWTMiddleware) {
	apiv1 := r.Group("/api")
	apiv1.Use(authMiddleware.MiddlewareFunc())
	{
		apiHandler := api.NewSysApiHandler(factoryIns)

		apiv1.POST("/add", apiHandler.Create)

		apiv1.DELETE("", apiHandler.DeleteBatch)

		apiv1.PATCH("", apiHandler.Update)

		apiv1.POST("/page", apiHandler.GetPage)
	}
}
