package menu

import (
	srvv1 "go-web/internal/admin/service/v1"
	"go-web/internal/admin/store"

	"github.com/casbin/casbin/v2"
)

type SysMenuHandler struct {
	srv     srvv1.Service
	factory store.Factory
}

func NewSysMenuHandler(factory store.Factory, enforcer *casbin.Enforcer) *SysMenuHandler {
	return &SysMenuHandler{
		srv:     srvv1.NewService(factory, enforcer),
		factory: factory,
	}
}
