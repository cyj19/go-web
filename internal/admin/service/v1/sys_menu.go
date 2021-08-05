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
	GetList(menu *model.SysMenu) ([]model.SysMenu, error)
	GetPage(menuPage *model.SysMenuPage) ([]model.SysMenu, int64, error)
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

func (m *menuService) GetList(menu *model.SysMenu) ([]model.SysMenu, error) {
	whereOrder := createSysMenuQueryCondition(menu)
	return m.factory.SysMenu().GetList(whereOrder...)
}

func (m *menuService) GetPage(menuPage *model.SysMenuPage) ([]model.SysMenu, int64, error) {
	whereOrder := createSysMenuQueryCondition(&menuPage.SysMenu)
	pageIndex := menuPage.PageIndex
	pageSize := menuPage.PageSize
	if pageIndex < 1 {
		pageIndex = 1
	}
	return m.factory.SysMenu().GetPage(pageIndex, pageSize, whereOrder...)
}

func createSysMenuQueryCondition(param *model.SysMenu) []model.WhereOrder {
	var whereOrder []model.WhereOrder
	if param != nil {
		if param.Name != "" {
			v := "%" + param.Name + "%"
			whereOrder = append(whereOrder, model.WhereOrder{Where: "name like ?", Value: []interface{}{v}})
		}

		whereOrder = append(whereOrder, model.WhereOrder{Where: "status = ?", Value: []interface{}{param.Status}})

	}

	return whereOrder
}
