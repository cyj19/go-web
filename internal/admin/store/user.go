package store

import "go-web/internal/pkg/model"

//UserStore defines the user storage interface.
type SysUserStore interface {
	Update(u *model.SysUser) error
	Delete(id uint64) error
	DeleteBatch(ids []uint64) error
	GetByUsername(username string) (*model.SysUser, error)
	List(whereOrder ...model.WhereOrder) ([]model.SysUser, error)
	GetPage(pageIndex int, pageSize int, whereOrder ...model.WhereOrder) ([]model.SysUser, int64, error)
	Login(username, password string) (*model.SysUser, error)
}
