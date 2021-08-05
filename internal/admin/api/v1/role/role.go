package role

import (
	srvv1 "go-web/internal/admin/service/v1"
	"go-web/internal/admin/store"

	"github.com/casbin/casbin/v2"
)

type SysRoleHandler struct {
	srv     srvv1.Service
	factory store.Factory
}

func NewSysRoleHandler(factory store.Factory, enforcer *casbin.Enforcer) *SysRoleHandler {
	return &SysRoleHandler{
		srv:     srvv1.NewService(factory, enforcer),
		factory: factory,
	}
}
