package role

import (
	srvv1 "github.com/vagaryer/go-web/internal/admin/service/v1"
	"github.com/vagaryer/go-web/internal/admin/store"
)

type SysRoleHandler struct {
	srv     srvv1.Service
	factory store.Factory
}

func NewSysRoleHandler(factory store.Factory) *SysRoleHandler {
	return &SysRoleHandler{
		srv:     srvv1.NewService(factory),
		factory: factory,
	}
}
