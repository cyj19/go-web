package router

import (
	"go-web/internal/admin/api/v1/role"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// 注册角色路由
func InitRoleRouter(r *gin.RouterGroup, factoryIns store.Factory) {

	rolev1 := r.Group("/role")
	// 使用casbin中间件
	rolev1.Use(middleware.CasbinMiddleware(initialize.GetEnforcerIns()))
	{
		roleHandler := role.NewRoleHandler(factoryIns)

		rolev1.GET("/:id", roleHandler.GetById)

		rolev1.POST("/add", roleHandler.Create)

		rolev1.POST("/permission", roleHandler.SetRolePermission)

		rolev1.DELETE("/:id", roleHandler.Delete)

		rolev1.DELETE("", roleHandler.DeleteBatch)

		rolev1.PUT("/update", roleHandler.Update)

		rolev1.POST("/list", roleHandler.List)

		rolev1.POST("/page", roleHandler.GetPage)

	}
}
