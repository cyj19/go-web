package store

import "go-web/internal/pkg/model"

//UserStore defines the user storage interface.
type SysUserStore interface {
	Update(u *model.SysUser) error
	UpdateRoleForUser(cd *model.CreateDelete) error
	BatchDelete(ids []uint64) error
	GetByUsername(username string) (*model.SysUser, error)
	Login(username, password string) (*model.SysUser, error)
}
