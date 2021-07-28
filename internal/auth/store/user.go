package store

import "go-web/internal/pkg/model"

type UserStore interface {
	Login(username, password string) (*model.SysUser, error)
}
