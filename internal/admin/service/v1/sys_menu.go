package v1

import (
	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"
)

type SysMenuSrv interface {
	Create(menu *model.SysMenu) error
	Update(menu *model.SysMenu) error
	BatchDelete(ids []uint64) error
	GetById(id uint64) (*model.SysMenu, error)
	GetByPath(path string) (*model.SysMenu, error)
	GetSome(ids []uint64) ([]model.SysMenu, error)
	GetList(whereOrders ...model.WhereOrder) ([]model.SysMenu, error)
	GetPage(pageIndex int, pageSize int, whereOrders ...model.WhereOrder) ([]model.SysMenu, int64, error)
}

type menuService struct {
	factory store.Factory
}

func newSysMenu(srv *service) SysMenuSrv {
	return &menuService{factory: srv.factory}
}

func (m *menuService) Create(menu *model.SysMenu) error {
	return m.factory.SysMenu().Create(menu)
}

func (m *menuService) Update(menu *model.SysMenu) error {
	return m.factory.SysMenu().Update(menu)
}

func (m *menuService) BatchDelete(ids []uint64) error {
	return m.factory.SysMenu().BatchDelete(ids)
}

func (m *menuService) GetById(id uint64) (*model.SysMenu, error) {
	return m.factory.SysMenu().GetById(id)
}

func (m *menuService) GetByPath(path string) (*model.SysMenu, error) {
	return m.factory.SysMenu().GetByPath(path)
}

func (m *menuService) GetSome(ids []uint64) ([]model.SysMenu, error) {
	return m.factory.SysMenu().GetSome(ids)
}

func (m *menuService) GetList(whereOrders ...model.WhereOrder) ([]model.SysMenu, error) {

	return m.factory.SysMenu().GetList(whereOrders...)
}

func (m *menuService) GetPage(pageIndex int, pageSize int, whereOrders ...model.WhereOrder) ([]model.SysMenu, int64, error) {
	if pageIndex <= 0 {
		pageIndex = 1
	}
	return m.factory.SysMenu().GetPage(pageIndex, pageSize, whereOrders...)
}
