package store

import "go-web/internal/pkg/model"

type SysRoleStore interface {
	Create(r *model.SysRole) error
	Update(r *model.SysRole) error
	Delete(id uint64) error
	DeleteBatch(ids []uint64) error
	GetById(id uint64) (*model.SysRole, error)
	List(whereOrder ...model.WhereOrder) ([]model.SysRole, error)
	GetPage(pageIndex int, pageSize int, whereOrder ...model.WhereOrder) ([]model.SysRole, int64, error)
}
