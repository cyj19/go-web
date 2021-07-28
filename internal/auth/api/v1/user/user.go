package user

import (
	srvv1 "go-web/internal/auth/service/v1"
	"go-web/internal/auth/store"
)

type UserHandler struct {
	srv     srvv1.Service
	factory store.Factory
}

func NewUserHandler(factory store.Factory) *UserHandler {
	return &UserHandler{srv: srvv1.NewService(factory), factory: factory}
}
