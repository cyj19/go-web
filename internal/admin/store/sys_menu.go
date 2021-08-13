package store

import "go-web/internal/pkg/model"

type SysMenuStore interface {
	GetByPath(path string) (*model.SysMenu, error)
	GetSome(ids []uint64) ([]model.SysMenu, error)
}
