package v1

import (
	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"
)

type SysApiSrv interface {
	Update(a *model.SysApi) error
	DeleteBatch(ids []uint64) error
	GetById(id uint64) (*model.SysApi, error)
	GetList(a *model.SysApi) ([]model.SysApi, error)
	GetPage(apiPage *model.SysApiPage) ([]model.SysApi, int64, error)
}

type apiService struct {
	factory store.Factory
}

func newSysApi(srv *service) SysApiSrv {
	return &apiService{factory: srv.factory}
}

func (a *apiService) Update(api *model.SysApi) error {
	return a.factory.SysApi().Update(api)
}

func (a *apiService) DeleteBatch(ids []uint64) error {
	return a.factory.SysApi().DeleteBatch(ids)
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
