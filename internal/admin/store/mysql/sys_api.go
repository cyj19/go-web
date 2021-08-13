package mysql

import (
	"go-web/internal/admin/store"

	"gorm.io/gorm"
)

type sysApi struct {
	db *gorm.DB
}

func newSysApi(ds *datastore) store.SysApiStore {
	return &sysApi{db: ds.db}
}

// func (a *sysApi) GetList(whereOrder ...model.WhereOrder) ([]model.SysApi, error) {
// 	result := make([]model.SysApi, 0)
// 	dbTemp := queryByCondition(a.db, &model.SysApi{}, whereOrder)
// 	err := dbTemp.Find(&result).Error
// 	return result, err
// }

// func (a *sysApi) GetPage(pageIndex int, pageSize int, whereOrder ...model.WhereOrder) ([]model.SysApi, int64, error) {
// 	result := make([]model.SysApi, 0)
// 	dbTemp := queryByCondition(a.db, &model.SysApi{}, whereOrder)

// 	// 查询总数
// 	var count int64
// 	var err error
// 	err = dbTemp.Count(&count).Error
// 	if err != nil || count == 0 {
// 		return nil, count, err
// 	}
// 	err = dbTemp.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&result).Error
// 	return result, count, err

// }
