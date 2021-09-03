package router

import (
	"github.com/vagaryer/go-web/internal/admin/api/v1/menu"
	"github.com/vagaryer/go-web/internal/admin/global"
	"github.com/vagaryer/go-web/internal/admin/store"
	"github.com/vagaryer/go-web/internal/pkg/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 注册菜单路由
func InitMenuRouter(r *gin.RouterGroup, factoryIns store.Factory, authMiddleware *jwt.GinJWTMiddleware) {

	menuv1 := r.Group("/menu")
	menuv1.Use(authMiddleware.MiddlewareFunc(), middleware.CasbinMiddleware(factoryIns, global.Conf, global.Enforcer))
	{
		menuHandler := menu.NewSysMenuHandler(factoryIns)

		menuv1.POST("/add", menuHandler.Create)
		menuv1.DELETE("/delete", menuHandler.BatchDelete)
		menuv1.PATCH("/update", menuHandler.Update)
		menuv1.POST("/list", menuHandler.GetList)
		menuv1.GET("/all", menuHandler.GetMenusByRoleId)
		menuv1.POST("/page", menuHandler.GetPage)
	}
}
