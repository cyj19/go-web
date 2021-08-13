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
