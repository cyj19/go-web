package v1

import (
	"go-web/internal/auth/store"
	"go-web/internal/pkg/model"
)

type UserSrv interface {
	Login(username, password string) (*model.SysUser, error)
}

type userService struct {
	factory store.Factory
}

func newUserService(s *service) UserSrv {
	return &userService{factory: s.factory}
}

func (u *userService) Login(username, password string) (*model.SysUser, error) {
	return u.factory.User().Login(username, password)
}
