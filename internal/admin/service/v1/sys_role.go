package v1

import (
	"context"
	"fmt"

	"github.com/vagaryer/go-web/internal/admin/global"
	"github.com/vagaryer/go-web/internal/admin/store"
	"github.com/vagaryer/go-web/internal/pkg/cache"
	"github.com/vagaryer/go-web/internal/pkg/model"
	"github.com/vagaryer/go-web/internal/pkg/util"
)

type SysRoleSrv interface {
	Create(ctx context.Context, values ...model.SysRole) error
	Update(ctx context.Context, value *model.SysRole) error
	UpdateMenuForRole(ctx context.Context, cd *model.CreateDelete) error
	UpdateApiForRole(ctx context.Context, cd *model.CreateDelete) error
	BatchDelete(ctx context.Context, ids []uint64) error
	GetById(ctx context.Context, id uint64) (*model.SysRole, error)
	GetByName(ctx context.Context, name string) (*model.SysRole, error)
	GetList(ctx context.Context, role model.SysRole) ([]model.SysRole, error)
	GetListByWhereOrder(ctx context.Context, whereOrders ...model.WhereOrder) ([]model.SysRole, error)
	GetPage(ctx context.Context, rolePage model.SysRolePage) (*model.Page, error)
}

type roleService struct {
	factory store.Factory
}

func newSysRole(srv *service) SysRoleSrv {
	return &roleService{
		factory: srv.factory,
	}
}

func (r *roleService) Create(ctx context.Context, values ...model.SysRole) error {
	err := r.factory.Create(&values)
	if err != nil {
		return err
	}
	// 清空缓存
	cleanCache(values[0].TableName() + "*")
	return nil
}

func (r *roleService) Update(ctx context.Context, role *model.SysRole) error {
	err := r.factory.Update(role)
	if err != nil {
		return err
	}
	// 清空缓存
	cleanCache(role.TableName() + "*")
	return nil
}

func (r *roleService) UpdateMenuForRole(ctx context.Context, cd *model.CreateDelete) error {
	// 查询记录是否存在
	_, err := r.GetById(ctx, cd.Id)
	if err != nil {
		return fmt.Errorf("记录不存在：%v ", err)
	}
	return r.factory.SysRole().UpdateMenuForRole(cd)
}

// 更新角色的接口权限，维护casbin规则
func (r *roleService) UpdateApiForRole(ctx context.Context, cd *model.CreateDelete) error {
	// 查询记录是否存在
	_, err := r.GetById(ctx, cd.Id)
	if err != nil {
		return fmt.Errorf("记录不存在：%v ", err)
	}
	// 创建api服务
	as := &apiService{factory: r.factory}
	// 创建casbin服务
	cs := &casbinService{factory: r.factory}
	// 删除接口权限
	if len(cd.Delete) > 0 {
		// 获取要删除的api
		whereOrder := model.WhereOrder{Where: "id in ?", Value: []interface{}{cd.Delete}}
		deleteApis, _ := as.GetListByWhereOrder(ctx, whereOrder)
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
			_, err = cs.BatchDeleteRoleCasbins(ctx, deleteCasbins)
			if err != nil {
				return err
			}
		}
	}

	// 增加接口权限
	if len(cd.Create) > 0 {
		// 获取要增加的api
		whereOrder := model.WhereOrder{Where: "id in ?", Value: []interface{}{cd.Create}}
		createApis, _ := as.GetListByWhereOrder(ctx, whereOrder)
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
			_, err = cs.BatchCreateRoleCasbins(ctx, createCasbins)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *roleService) BatchDelete(ctx context.Context, ids []uint64) error {
	value := new(model.SysRole)
	err := r.factory.BatchDelete(ids, value)
	if err != nil {
		return err
	}
	// 清空缓存
	cleanCache(value.TableName() + "*")
	return nil
}

func (r *roleService) GetById(ctx context.Context, id uint64) (*model.SysRole, error) {
	value := new(model.SysRole)
	key := fmt.Sprintf("%s:id:%d", value.TableName(), id)
	err := cache.Get(global.RedisIns, key, value)
	if err != nil {
		err = r.factory.GetById(id, value)
		// 写入缓存
		cache.Set(global.RedisIns, key, value)

	}
	return value, err
}

func (r *roleService) GetByName(ctx context.Context, name string) (*model.SysRole, error) {
	value := new(model.SysRole)
	key := fmt.Sprintf("%s:name:%s", value.TableName(), name)
	err := cache.Get(global.RedisIns, key, value)
	if err != nil {
		value, err = r.factory.SysRole().GetByName(name)
		// 写入缓存
		cache.Set(global.RedisIns, key, value)
	}
	return value, err
}

func (r *roleService) GetList(ctx context.Context, role model.SysRole) ([]model.SysRole, error) {
	var list []model.SysRole
	var err error
	var key string
	key = fmt.Sprintf("%s:id:%d:name:%s:nameZh:%s", role.TableName(), role.Id, role.Name, role.NameZh)
	if role.Status != nil {
		key = fmt.Sprintf("%s:status:%t", key, *role.Status)
	}
	if role.Sort != nil {
		key = fmt.Sprintf("%s:sort:%d", key, *role.Sort)
	}

	list = cache.GetSysRoleList(global.RedisIns, key)
	if len(list) < 1 {
		whereOrders := util.GenWhereOrderByStruct(role)
		err = r.factory.GetList(&model.SysRole{}, &list, whereOrders...)
		// 添加到缓存
		cache.SetSysRoleList(global.RedisIns, key, list)
	}
	return list, err
}

// 特定条件的查询
func (r *roleService) GetListByWhereOrder(ctx context.Context, whereOrders ...model.WhereOrder) ([]model.SysRole, error) {
	list := make([]model.SysRole, 0)
	err := r.factory.GetList(&model.SysRole{}, &list, whereOrders...)
	return list, err
}

func (r *roleService) GetPage(ctx context.Context, rolePage model.SysRolePage) (*model.Page, error) {
	var list []model.SysRole
	var err error
	var key string
	var count int64
	pageIndex := rolePage.PageIndex
	pageSize := rolePage.PageSize
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = defaultSize
	}
	key = fmt.Sprintf("%s:id:%d:name:%s:nameZh:%s", rolePage.TableName(), rolePage.Id, rolePage.Name, rolePage.NameZh)
	if rolePage.Status != nil {
		key = fmt.Sprintf("%s:status:%t", key, *rolePage.Status)
	}
	if rolePage.Sort != nil {
		key = fmt.Sprintf("%s:sort:%d", key, *rolePage.Sort)
	}
	key = fmt.Sprintf("%s:pageIndex:%d:pageSize:%d", key, pageIndex, pageSize)

	// 从缓存中查找
	list = cache.GetSysRoleList(global.RedisIns, key)
	if len(list) < 1 {
		whereOrders := util.GenWhereOrderByStruct(rolePage.SysRole)
		count, err = r.factory.GetPage(pageIndex, pageSize, &model.SysRole{}, &list, whereOrders...)
		// 添加到缓存
		cache.SetSysRoleList(global.RedisIns, key, list)
	}

	page := &model.Page{
		Records:  list,
		Total:    count,
		PageInfo: model.PageInfo{PageIndex: pageIndex, PageSize: pageSize},
	}
	page.SetPageNum(count)
	return page, err
}
