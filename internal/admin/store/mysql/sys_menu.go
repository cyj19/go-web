package mysql

import (
	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"

	"gorm.io/gorm"
)

type menu struct {
	db *gorm.DB
}

func newSysMenu(ds *datastore) store.SysMenuStore {
	return &menu{db: ds.db}
}

//实现MenuStore接口

func (m *menu) Create(menu *model.SysMenu) error {
	return m.db.Create(menu).Error
}

func (m *menu) Update(menu *model.SysMenu) error {
	return m.db.Save(menu).Error
}

func (m *menu) BatchDelete(ids []uint64) error {
	return batchDelete(m.db, &model.SysMenu{}, ids)
}

// func (m *menu) GetById(id uint64) (*model.SysMenu, error) {
// 	var result model.SysMenu
// 	err := m.db.Preload("Roles").Where("id = ?", id).First(&result).Error
// 	return &result, err
// }

func (m *menu) GetByPath(path string) (*model.SysMenu, error) {
	var result model.SysMenu
	err := m.db.Preload("Roles").Where("path = ?", path).First(&result).Error
	return &result, err
}

func (m *menu) GetSome(ids []uint64) ([]model.SysMenu, error) {
	var result []model.SysMenu
	err := m.db.Where("id in (?)", ids).Find(&result).Error
	return result, err
}

// func (m *menu) GetList(whereOrder ...model.WhereOrder) ([]model.SysMenu, error) {
// 	var result []model.SysMenu
// 	tx := queryByCondition(m.db, &model.SysMenu{}, whereOrder)
// 	err := tx.Find(&result).Error
// 	return result, err
// }

// func (m *menu) GetPage(pageIndex int, pageSize int, whereOrder ...model.WhereOrder) ([]model.SysMenu, int64, error) {
// 	var result []model.SysMenu
// 	tx := queryByCondition(m.db, &model.SysMenu{}, whereOrder)
// 	//查询总数
// 	var count int64
// 	var err error
// 	err = tx.Count(&count).Error
// 	if err != nil || count == 0 {
// 		return nil, count, err
// 	}
// 	err = tx.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&result).Error
// 	return result, count, err
// }
