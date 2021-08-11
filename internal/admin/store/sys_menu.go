package store

import "go-web/internal/pkg/model"

type SysMenuStore interface {
	Create(menu *model.SysMenu) error
	Update(menu *model.SysMenu) error
	BatchDelete(ids []uint64) error
	GetByPath(path string) (*model.SysMenu, error)
	GetSome(ids []uint64) ([]model.SysMenu, error)
}
