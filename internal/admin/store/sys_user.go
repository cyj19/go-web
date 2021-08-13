package store

import "go-web/internal/pkg/model"

//UserStore defines the user storage interface.
type SysUserStore interface {
	UpdateRoleForUser(cd *model.CreateDelete) error
	GetByUsername(username string) (*model.SysUser, error)
	Login(username, password string) (*model.SysUser, error)
}
