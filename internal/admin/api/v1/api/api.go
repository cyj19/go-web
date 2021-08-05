package api

import (
	srvv1 "go-web/internal/admin/service/v1"
	"go-web/internal/admin/store"

	"github.com/casbin/casbin/v2"
)

type SysApiHandler struct {
	srv     srvv1.Service
	factory store.Factory
}

func NewSysApiHandler(factory store.Factory, enforcer *casbin.Enforcer) *SysApiHandler {
	return &SysApiHandler{
		srv:     srvv1.NewService(factory, enforcer),
		factory: factory,
	}
}
