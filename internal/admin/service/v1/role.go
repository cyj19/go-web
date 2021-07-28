package v1

import (
	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"
)

type SysRoleSrv interface {
	Create(r *model.SysRole) error
	Update(r *model.SysRole) error
	Delete(id uint64) error
	DeleteBatch(ids []uint64) error
	GetById(id uint64) (*model.SysRole, error)
	List(r *model.SysRole) ([]model.SysRole, error)
	GetPage(rolePage *model.SysRolePage) ([]model.SysRole, int64, error)
}

type roleService struct {
	factory store.Factory
}

func newSysRole(srv *service) SysRoleSrv {
	return &roleService{factory: srv.factory}
}

func (r *roleService) Create(role *model.SysRole) error {
	return r.factory.SysRole().Create(role)
}

func (r *roleService) Update(role *model.SysRole) error {
	return r.factory.SysRole().Update(role)
}

func (r *roleService) Delete(id uint64) error {
	return r.factory.SysRole().Delete(id)
}

func (r *roleService) DeleteBatch(ids []uint64) error {
	return r.factory.SysRole().DeleteBatch(ids)
}

func (r *roleService) GetById(id uint64) (*model.SysRole, error) {
	return r.factory.SysRole().GetById(id)
}

func (r *roleService) List(role *model.SysRole) ([]model.SysRole, error) {
	whereOrder := createSysRoleCondition(role)
	return r.factory.SysRole().List(whereOrder...)
}

func (r *roleService) GetPage(userPage *model.SysRolePage) ([]model.SysRole, int64, error) {
	whereOrder := createSysRoleCondition(&userPage.SysRole)
	pageIndex := userPage.PageIndex
	pageSize := userPage.PageSize
	if pageIndex < 1 {
		pageIndex = 1
	}
	return r.factory.SysRole().GetPage(pageIndex, pageSize, whereOrder...)
}

func createSysRoleCondition(param *model.SysRole) []model.WhereOrder {
	var whereOrder []model.WhereOrder
	if param != nil {
		if param.Name != "" {
			v := "%" + param.Name + "%"
			whereOrder = append(whereOrder, model.WhereOrder{Where: "name like ?", Value: []interface{}{v}})
		}
		if param.NameZh != "" {
			v := "%" + param.NameZh + "%"
			whereOrder = append(whereOrder, model.WhereOrder{Where: "name_zh like ?", Value: []interface{}{v}})
		}
	}

	return whereOrder
}
