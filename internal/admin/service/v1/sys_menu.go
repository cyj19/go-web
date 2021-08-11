package v1

import (
	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"
)

type SysMenuSrv interface {
	Create(menu *model.SysMenu) error
	Update(menu *model.SysMenu) error
	BatchDelete(ids []uint64) error
	GetByPath(path string) (*model.SysMenu, error)
	GetSome(ids []uint64) ([]model.SysMenu, error)
	GetList(whereOrders ...model.WhereOrder) ([]model.SysMenu, error)
	GetMenusByRoleId(ids []uint64) ([]model.SysMenu, error)
	GetPage(pageIndex int, pageSize int, whereOrders ...model.WhereOrder) (*model.Page, error)
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

func (m *menuService) GetByPath(path string) (*model.SysMenu, error) {
	return m.factory.SysMenu().GetByPath(path)
}

func (m *menuService) GetSome(ids []uint64) ([]model.SysMenu, error) {
	return m.factory.SysMenu().GetSome(ids)
}

func (m *menuService) GetList(whereOrders ...model.WhereOrder) ([]model.SysMenu, error) {
	list := make([]model.SysMenu, 0)
	err := m.factory.GetList(model.SysMenu{}, &list, whereOrders...)
	return list, err
}

func (m *menuService) GetMenusByRoleId(ids []uint64) ([]model.SysMenu, error) {
	// 创建role服务
	rs := &roleService{factory: m.factory}
	whereOrder := model.WhereOrder{Where: "id in ?", Value: []interface{}{ids}}
	roles, err := rs.GetListByWhereOrder(whereOrder)
	if err != nil {
		return nil, err
	}
	// 角色拥有的菜单
	menus := make([]model.SysMenu, 0)
	for i, role := range roles {
		if i > 0 {
			// 已有的不加入
			for _, menu := range role.Menus {
				// 判断菜单是否已存在
				if !util.ContainsSysMenu(menus, menu) {
					menus = append(menus, menu)
				}
			}
		} else {
			menus = append(menus, role.Menus...)
		}

	}
	// 根据parentId 和 sort构建菜单树
	tree := genMenuTree(0, menus)
	return tree, nil
}

func (m *menuService) GetPage(pageIndex int, pageSize int, whereOrders ...model.WhereOrder) (*model.Page, error) {
	list := make([]model.SysMenu, 0)
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = defaultSize
	}
	count, err := m.factory.GetPage(pageIndex, pageSize, model.SysMenu{}, &list, whereOrders...)
	page := &model.Page{
		Records:  list,
		Total:    count,
		PageInfo: model.PageInfo{PageIndex: pageIndex, PageSize: pageSize},
	}
	page.SetPageNum(count)
	return page, err
}

func genMenuTree(parentId uint64, menus []model.SysMenu) []model.SysMenu {
	tree := make([]model.SysMenu, 0)
	for _, menu := range menus {
		if menu.ParentId == parentId {
			// 递归遍历子菜单
			menu.Children = genMenuTree(menu.Id, menus)
			tree = append(tree, menu)
		}
	}
	return tree
}
