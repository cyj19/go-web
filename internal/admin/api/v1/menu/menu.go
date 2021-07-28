package menu

import (
	srvv1 "go-web/internal/admin/service/v1"
	"go-web/internal/admin/store"
)

type MenuHandler struct {
	srv     srvv1.Service
	factory store.Factory
}

func NewMenuHandler(factory store.Factory) *MenuHandler {
	return &MenuHandler{
		srv:     srvv1.NewService(factory),
		factory: factory,
	}
}
