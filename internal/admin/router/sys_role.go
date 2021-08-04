package router

import (
	"go-web/internal/admin/api/v1/role"
	"go-web/internal/admin/store"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 注册角色路由
func InitRoleRouter(r *gin.RouterGroup, factoryIns store.Factory, authMiddleware *jwt.GinJWTMiddleware) {

	rolev1 := r.Group("/role")
	rolev1.Use(authMiddleware.MiddlewareFunc())
	// rolev1.Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(initialize.GetEnforcerIns()))
	{
		roleHandler := role.NewSysRoleHandler(factoryIns)

		rolev1.GET("/:id", roleHandler.GetById)

		rolev1.POST("/add", roleHandler.Create)

		rolev1.PATCH("/update", roleHandler.Update)

		rolev1.POST("/list", roleHandler.GetList)

		rolev1.POST("/page", roleHandler.GetPage)

		rolev1.DELETE("", roleHandler.DeleteBatch)
	}
}
