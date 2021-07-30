package router

import (
	"go-web/internal/admin/api/v1/menu"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// 注册菜单路由
func InitMenuRouter(r *gin.RouterGroup, factoryIns store.Factory) {

	menuv1 := r.Group("/menu")
	menuv1.Use(middleware.CasbinMiddleware(initialize.GetEnforcerIns()))
	{
		menuHandler := menu.NewMenuHandler(factoryIns)

		menuv1.POST("/add", menuHandler.Create)
		menuv1.DELETE("/:id", menuHandler.Delete)
		menuv1.DELETE("", menuHandler.DeleteBatch)
		menuv1.PUT("/update", menuHandler.Update)
		menuv1.POST("/list", menuHandler.List)
		menuv1.POST("/page", menuHandler.GetPage)
		menuv1.GET("/:id", menuHandler.GetById)
	}
}
