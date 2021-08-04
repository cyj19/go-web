package store

import "go-web/internal/pkg/model"

type SysRoleStore interface {
	Create(r *model.SysRole) error
	Update(r *model.SysRole) error
	DeleteBatch(ids []uint64) error
	GetById(id uint64) (*model.SysRole, error)
	GetByName(name string) (*model.SysRole, error)
	GetList(whereOrder ...model.WhereOrder) ([]model.SysRole, error)
	GetPage(pageIndex int, pageSize int, whereOrder ...model.WhereOrder) ([]model.SysRole, int64, error)
}
