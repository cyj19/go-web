package v1

import (
	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"
	"go-web/internal/pkg/util"

	"github.com/casbin/casbin/v2"
)

type SysApiSrv interface {
	Create(a *model.SysApi) error
	Update(a *model.SysApi) error
	BatchDelete(ids []uint64) error
	GetById(id uint64) (*model.SysApi, error)
	GetList(a *model.SysApi) ([]model.SysApi, error)
	GetPage(apiPage *model.SysApiPage) ([]model.SysApi, int64, error)
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
func (a *apiService) Create(api *model.SysApi) error {
	err := a.factory.Create(api)
	if err != nil {
		return err
	}

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

	return nil

}

func (a *apiService) Update(api *model.SysApi) error {
	// 查询接口是否存在
	oldApi, err := a.GetById(api.Id)
	if err != nil {
		return err
	}
	// 更新接口
	err = a.factory.SysApi().Update(api)
	if err != nil {
		return err
	}

	// 对比新旧接口的Method , Path
	if oldApi.Method != api.Method || oldApi.Path != api.Path {
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
					Path:    api.Path,
					Method:  api.Method,
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
	whereOrder := model.WhereOrder{Where: "id in ?", Value: []interface{}{ids}}
	apis, err := a.factory.SysApi().GetList(whereOrder)
	if err != nil {
		return err
	}
	//  删除接口
	err = a.factory.SysApi().BatchDelete(ids)
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
	return a.factory.SysApi().GetById(id)
}

func (a *apiService) GetList(api *model.SysApi) ([]model.SysApi, error) {
	whereOrders := createSysApiQueryCondition(api)
	return a.factory.SysApi().GetList(whereOrders...)
}

func (a *apiService) GetPage(apiPage *model.SysApiPage) ([]model.SysApi, int64, error) {
	whereOrders := createSysApiQueryCondition(&apiPage.SysApi)
	pageIndex := apiPage.PageIndex
	pageSize := apiPage.PageSize
	if pageIndex <= 0 {
		pageIndex = 1
	}
	return a.factory.SysApi().GetPage(pageIndex, pageSize, whereOrders...)
}

func createSysApiQueryCondition(param *model.SysApi) []model.WhereOrder {
	if param != nil {
		whereOrders := make([]model.WhereOrder, 0)
		if param.Id > 0 {
			whereOrders = append(whereOrders, model.WhereOrder{Where: "id = ?", Value: []interface{}{param.Id}})
		}
		if param.Method != "" {
			v := "%" + param.Method + "%"
			whereOrders = append(whereOrders, model.WhereOrder{Where: "method like ?", Value: []interface{}{v}})
		}
		if param.Path != "" {
			v := "%" + param.Path + "%"
			whereOrders = append(whereOrders, model.WhereOrder{Where: "path like ?", Value: []interface{}{v}})
		}
		if param.Category != "" {
			v := "%" + param.Category + "%"
			whereOrders = append(whereOrders, model.WhereOrder{Where: "category like ?", Value: []interface{}{v}})
		}
		if param.Creator != "" {
			v := "%" + param.Creator + "%"
			whereOrders = append(whereOrders, model.WhereOrder{Where: "creator like ?", Value: []interface{}{v}})
		}
		return whereOrders
	}
	return nil
}
