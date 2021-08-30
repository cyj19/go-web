package router

import (
	"go-web/internal/admin/api/v1/role"
	"go-web/internal/admin/global"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 注册角色路由
func InitRoleRouter(r *gin.RouterGroup, factoryIns store.Factory, authMiddleware *jwt.GinJWTMiddleware) {

	rolev1 := r.Group("/role")
	rolev1.Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(factoryIns, global.Conf, global.Enforcer))
	{
		roleHandler := role.NewSysRoleHandler(factoryIns)

		rolev1.POST("/add", roleHandler.Create)
		rolev1.DELETE("/delete", roleHandler.BatchDelete)
		rolev1.PATCH("/update", roleHandler.Update)
		rolev1.PATCH("/menu/update", roleHandler.UpdateMenuForRole)
		rolev1.PATCH("/api/update", roleHandler.UpdateApiForRole)
		rolev1.POST("/list", roleHandler.GetList)
		rolev1.POST("/page", roleHandler.GetPage)
		rolev1.GET("/:id", roleHandler.GetById)
	}
}
