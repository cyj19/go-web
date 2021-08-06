package router

import (
	"go-web/internal/admin/api/v1/menu"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// 注册菜单路由
func InitMenuRouter(r *gin.RouterGroup, factoryIns store.Factory, enforcer *casbin.Enforcer, authMiddleware *jwt.GinJWTMiddleware) {

	menuv1 := r.Group("/menu")
	menuv1.Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(factoryIns, enforcer))
	{
		menuHandler := menu.NewSysMenuHandler(factoryIns, enforcer)

		menuv1.POST("/add", menuHandler.Create)
		menuv1.DELETE("/delete", menuHandler.BatchDelete)
		menuv1.PUT("/update", menuHandler.Update)
		menuv1.POST("/list", menuHandler.GetList)
		menuv1.POST("/page", menuHandler.GetPage)
	}
}
