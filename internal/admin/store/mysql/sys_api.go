package mysql

import (
	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"

	"gorm.io/gorm"
)

type sysApi struct {
	db *gorm.DB
}

func newSysApi(ds *datastore) store.SysApiStore {
	return &sysApi{db: ds.db}
}

func (a *sysApi) Update(api *model.SysApi) error {
	return a.db.Updates(api).Error
}

func (a *sysApi) DeleteBatch(ids []uint64) error {
	return deleteBatch(a.db, &model.SysApi{}, ids)
}

func (a *sysApi) GetById(id uint64) (*model.SysApi, error) {
	api := new(model.SysApi)
	err := a.db.Where("id = ?", id).First(api).Error
	return api, err
}

func (a *sysApi) GetList(whereOrder ...model.WhereOrder) ([]model.SysApi, error) {
	result := make([]model.SysApi, 0)
	dbTemp := queryByCondition(a.db, &model.SysApi{}, whereOrder)
	err := dbTemp.Find(result).Error
	return result, err
}

func (a *sysApi) GetPage(pageIndex int, pageSize int, whereOrder ...model.WhereOrder) ([]model.SysApi, int64, error) {
	result := make([]model.SysApi, 0)
	dbTemp := queryByCondition(a.db, &model.SysApi{}, whereOrder)

	// 查询总数
	var count int64
	var err error
	err = dbTemp.Count(&count).Error
	if err != nil || count == 0 {
		return nil, count, err
	}
	err = dbTemp.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(result).Error
	return result, count, err

}
