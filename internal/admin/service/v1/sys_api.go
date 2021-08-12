package v1

import (
	"fmt"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/cache"
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"

	"github.com/casbin/casbin/v2"
)

type SysApiSrv interface {
	Create(values ...model.SysApi) error
	Update(value *model.SysApi) error
	BatchDelete(ids []uint64) error
	GetById(id uint64) (*model.SysApi, error)
	GetList(whereOrders ...model.WhereOrder) ([]model.SysApi, error)
	GetPage(pageIndex int, pageSize int, whereOrders ...model.WhereOrder) (*model.Page, error)
}

type apiService struct {
	factory  store.Factory
	enforcer *casbin.Enforcer
}

func newSysApi(srv *service) SysApiSrv {
	return &apiService{
		factory:  srv.factory,
		enforcer: srv.enforcer,
	}
}

// 自定义接口创建，同步创建casbin规则
func (a *apiService) Create(values ...model.SysApi) error {
	err := a.factory.Create(&values)
	if err != nil {
		return err
	}

	for _, api := range values {
		if len(api.Roles) > 0 {
			// 创建casbin规则
			cs := &casbinService{enforcer: a.enforcer}
			roleCasbins := make([]model.SysRoleCasbin, 0)
			for _, role := range api.Roles {
				roleCasbins = append(roleCasbins, model.SysRoleCasbin{
					Kyeword: util.Uint642Str(role),
					Path:    api.Path,
					Method:  api.Method,
				})
			}
			_, err := cs.BatchCreateRoleCasbins(roleCasbins)
			return err
		}
	}

	return nil

}

func (a *apiService) Update(value *model.SysApi) error {
	// 查询接口是否存在
	oldApi, err := a.GetById(value.Id)
	if err != nil {
		return err
	}
	// 更新接口
	err = a.factory.Update(value)
	if err != nil {
		return err
	}

	// 对比新旧接口的Method , Path
	if oldApi.Method != value.Method || oldApi.Path != value.Path {
		// 有修改，更新casbin规则
		cs := &casbinService{enforcer: a.enforcer}
		// 获取和旧接口相关的规则
		oldRules := cs.GetRoleCasbins(model.SysRoleCasbin{Path: oldApi.Path, Method: oldApi.Method})
		if len(oldRules) > 0 {
			// 删除旧规则
			cs.BatchDeleteRoleCasbins(oldRules)
			// 创建新规则
			newRules := make([]model.SysRoleCasbin, 0)
			for _, rule := range oldRules {
				newRules = append(newRules, model.SysRoleCasbin{
					Kyeword: rule.Kyeword,
					Path:    value.Path,
					Method:  value.Method,
				})
			}
			_, err = cs.BatchCreateRoleCasbins(newRules)
			return err
		}

	}

	return nil
}

func (a *apiService) BatchDelete(ids []uint64) error {
	// 查询接口
	apis := make([]model.SysApi, 0)
	whereOrder := model.WhereOrder{Where: "id in ?", Value: []interface{}{ids}}
	err := a.factory.GetList(model.SysApi{}, &apis, whereOrder)
	if err != nil {
		return err
	}
	//  删除接口
	err = a.factory.BatchDelete(ids, &model.SysApi{})
	if err != nil {
		return err
	}

	// 删除casbin规则
	cs := &casbinService{enforcer: a.enforcer}
	roleCasbins := make([]model.SysRoleCasbin, 0)
	for _, api := range apis {
		roleCasbins = append(roleCasbins, model.SysRoleCasbin{
			Path:   api.Path,
			Method: api.Method,
		})
	}
	_, err = cs.BatchDeleteRoleCasbins(roleCasbins)
	return err
}

func (a *apiService) GetById(id uint64) (*model.SysApi, error) {
	value := new(model.SysApi)
	key := fmt.Sprintf("%s:id:%d", value.TableName(), id)
	err := cache.Get(key, value)
	if err != nil {
		err = a.factory.GetById(id, value)
		// 写入缓存
		cache.Set(key, value)

	}
	return value, err
}

func (a *apiService) GetList(whereOrders ...model.WhereOrder) ([]model.SysApi, error) {
	list := make([]model.SysApi, 0)
	err := a.factory.GetList(model.SysApi{}, &list, whereOrders...)
	return list, err
}

func (a *apiService) GetPage(pageIndex int, pageSize int, whereOrders ...model.WhereOrder) (*model.Page, error) {
	list := make([]model.SysApi, 0)
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = defaultSize
	}
	count, err := a.factory.GetPage(pageIndex, pageSize, model.SysApi{}, &list, whereOrders...)
	page := &model.Page{
		Records:  list,
		Total:    count,
		PageInfo: model.PageInfo{PageIndex: pageIndex, PageSize: pageSize},
	}
	page.SetPageNum(count)
	return page, err
}
