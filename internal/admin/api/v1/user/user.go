package user

import (
	srvv1 "go-web/internal/admin/service/v1"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type SysUserHandler struct {
	srv     srvv1.Service
	factory store.Factory
}

func NewSysUserHandler(factory store.Factory, enforcer *casbin.Enforcer) *SysUserHandler {
	return &SysUserHandler{
		srv:     srvv1.NewService(factory, enforcer),
		factory: factory,
	}
}

func GetCurrentUser(c *gin.Context, factory store.Factory, enforcer *casbin.Enforcer) *model.SysUser {
	userId := c.MustGet("user")
	// 查询用户
	srv := srvv1.NewService(factory, enforcer)
	user, _ := srv.SysUser().GetById(userId.(uint64))
	return user
}
