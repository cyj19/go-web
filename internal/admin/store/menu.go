package store

import "go-web/internal/pkg/model"

type SysMenuStore interface {
	Create(menu *model.SysMenu) error
	Update(menu *model.SysMenu) error
	Delete(id uint64) error
	DeleteBatch(ids []uint64) error
	GetById(id uint64) (*model.SysMenu, error)
	GetSome(ids []uint64) ([]model.SysMenu, error)
	List(whereOrder ...model.WhereOrder) ([]model.SysMenu, error)
	GetPage(pageIndex int, pageSize int, whereOrder ...model.WhereOrder) ([]model.SysMenu, int64, error)
}
