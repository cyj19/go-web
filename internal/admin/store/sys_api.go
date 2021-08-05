package store

import "go-web/internal/pkg/model"

type SysApiStore interface {
	Create(a *model.SysApi) error
	Update(a *model.SysApi) error
	BatchDelete(ids []uint64) error
	GetById(id uint64) (*model.SysApi, error)
	GetList(whereOrder ...model.WhereOrder) ([]model.SysApi, error)
	GetPage(pageIndex int, pageSize int, whereOrder ...model.WhereOrder) ([]model.SysApi, int64, error)
}
