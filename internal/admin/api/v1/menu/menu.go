package menu

import (
	srvv1 "github.com/vagaryer/go-web/internal/admin/service/v1"
	"github.com/vagaryer/go-web/internal/admin/store"
)

type SysMenuHandler struct {
	srv     srvv1.Service
	factory store.Factory
}

func NewSysMenuHandler(factory store.Factory) *SysMenuHandler {
	return &SysMenuHandler{
		srv:     srvv1.NewService(factory),
		factory: factory,
	}
}
