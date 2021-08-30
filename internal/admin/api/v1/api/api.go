package api

import (
	srvv1 "go-web/internal/admin/service/v1"
	"go-web/internal/admin/store"
)

type SysApiHandler struct {
	srv     srvv1.Service
	factory store.Factory
}

func NewSysApiHandler(factory store.Factory) *SysApiHandler {
	return &SysApiHandler{
		srv:     srvv1.NewService(factory),
		factory: factory,
	}
}
