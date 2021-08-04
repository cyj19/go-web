package store

import "go-web/internal/pkg/model"

//UserStore defines the user storage interface.
type SysUserStore interface {
	Update(u *model.SysUser) error
	UpdateRoleForUser(cd *model.CreateDelete) error
	DeleteBatch(ids []uint64) error
	GetById(id uint64) (*model.SysUser, error)
	GetByUsername(username string) (*model.SysUser, error)
	GetList(whereOrder ...model.WhereOrder) ([]model.SysUser, error)
	GetPage(pageIndex int, pageSize int, whereOrder ...model.WhereOrder) ([]model.SysUser, int64, error)
	Login(username, password string) (*model.SysUser, error)
}
