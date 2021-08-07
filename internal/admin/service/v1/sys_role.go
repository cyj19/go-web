package v1

import (
	"fmt"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"

	"github.com/casbin/casbin/v2"
)

type SysRoleSrv interface {
	Create(r *model.SysRole) error
	Update(r *model.SysRole) error
	UpdateMenuForRole(cd *model.CreateDelete) error
	UpdateApiForRole(cd *model.CreateDelete) error
	BatchDelete(ids []uint64) error
	GetById(id uint64) (*model.SysRole, error)
	GetByName(name string) (*model.SysRole, error)
	GetList(whereOrders ...model.WhereOrder) ([]model.SysRole, error)
	GetPage(pageIndex int, pageSize int, whereOrders ...model.WhereOrder) (*model.Page, error)
}

type roleService struct {
	factory  store.Factory
	enforcer *casbin.Enforcer
}

func newSysRole(srv *service) SysRoleSrv {
	return &roleService{
		factory:  srv.factory,
		enforcer: srv.enforcer,
	}
}

func (r *roleService) Create(role *model.SysRole) error {
	return r.factory.SysRole().Create(role)
}

func (r *roleService) Update(role *model.SysRole) error {
	return r.factory.SysRole().Update(role)
}

func (r *roleService) UpdateMenuForRole(cd *model.CreateDelete) error {
	// 查询记录是否存在
	_, err := r.GetById(cd.Id)
	if err != nil {
		return fmt.Errorf("记录不存在：%v ", err)
	}
	return r.factory.SysRole().UpdateMenuForRole(cd)
}

// 更新角色的接口权限，维护casbin规则
func (r *roleService) UpdateApiForRole(cd *model.CreateDelete) error {
	// 查询记录是否存在
	_, err := r.GetById(cd.Id)
	if err != nil {
		return fmt.Errorf("记录不存在：%v ", err)
	}
	// 创建api服务
	as := &apiService{factory: r.factory, enforcer: r.enforcer}
	// 创建casbin服务
	cs := &casbinService{enforcer: r.enforcer}
	// 删除接口权限
	if len(cd.Delete) > 0 {
		// 获取要删除的api
		whereOrder := model.WhereOrder{Where: "id in ?", Value: []interface{}{cd.Delete}}
		deleteApis, _ := as.GetList(whereOrder)
		// 构建casbin规则
		deleteCasbins := make([]model.SysRoleCasbin, 0)
		for _, api := range deleteApis {
			deleteCasbins = append(deleteCasbins, model.SysRoleCasbin{
				Kyeword: util.Uint642Str(cd.Id),
				Path:    api.Path,
				Method:  api.Method,
			})
		}
		if len(deleteCasbins) > 0 {
			// 删除casbin规则
			_, err = cs.BatchDeleteRoleCasbins(deleteCasbins)
			if err != nil {
				return err
			}
		}
	}

	// 增加接口权限
	if len(cd.Create) > 0 {
		// 获取要增加的api
		whereOrder := model.WhereOrder{Where: "id in ?", Value: []interface{}{cd.Create}}
		createApis, _ := as.GetList(whereOrder)
		// 构建casbin规则
		createCasbins := make([]model.SysRoleCasbin, 0)
		for _, api := range createApis {
			createCasbins = append(createCasbins, model.SysRoleCasbin{
				Kyeword: util.Uint642Str(cd.Id),
				Path:    api.Path,
				Method:  api.Method,
			})
		}
		if len(createCasbins) > 0 {
			// 增加casbin规则
			_, err = cs.BatchCreateRoleCasbins(createCasbins)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *roleService) BatchDelete(ids []uint64) error {
	return r.factory.SysRole().BatchDelete(ids)
}

func (r *roleService) GetById(id uint64) (*model.SysRole, error) {
	return r.factory.SysRole().GetById(id)
}

func (r *roleService) GetByName(name string) (*model.SysRole, error) {
	return r.factory.SysRole().GetByName(name)
}

func (r *roleService) GetList(whereOrders ...model.WhereOrder) ([]model.SysRole, error) {
	return r.factory.SysRole().GetList(whereOrders...)
}

func (r *roleService) GetPage(pageIndex int, pageSize int, whereOrders ...model.WhereOrder) (*model.Page, error) {
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = defaultSize
	}
	list, count, err := r.factory.SysRole().GetPage(pageIndex, pageSize, whereOrders...)
	page := &model.Page{
		Records:  list,
		PageInfo: model.PageInfo{PageIndex: pageIndex, PageSize: pageSize},
	}
	page.SetPageNum(count)
	return page, err
}
