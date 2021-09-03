package user

import (
	srvv1 "github.com/vagaryer/go-web/internal/admin/service/v1"
	"github.com/vagaryer/go-web/internal/admin/store"
)

type SysUserHandler struct {
	srv     srvv1.Service
	factory store.Factory
}

func NewSysUserHandler(factory store.Factory) *SysUserHandler {
	return &SysUserHandler{
		srv:     srvv1.NewService(factory),
		factory: factory,
	}
}
