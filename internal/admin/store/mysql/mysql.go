package mysql

import (
	"fmt"
	"sync"

	"go-web/internal/admin/store"
	"go-web/internal/pkg/initialize"
	"go-web/internal/pkg/model"

	"gorm.io/gorm"
)

type datastore struct {
	db *gorm.DB
}

//实现Factory接口
func (ds *datastore) SysUser() store.SysUserStore {
	return newSysUser(ds)
}

func (ds *datastore) SysRole() store.SysRoleStore {
	return newSysRole(ds)
}

func (ds *datastore) SysMenu() store.SysMenuStore {
	return newSysMenu(ds)
}

func (ds *datastore) SysApi() store.SysApiStore {
	return newSysApi(ds)
}

func (ds *datastore) Create(value interface{}) error {
	return ds.db.Create(value).Error
}

//不能放到pkg包中
var (
	factory     store.Factory
	onceFactory sync.Once
)

//单例 获取factory
func GetMySQLFactory() (store.Factory, error) {
	db := initialize.GetMySQLIns()
	if db == nil {
		return nil, fmt.Errorf("mysql db is nil")
	}
	onceFactory.Do(func() {
		factory = &datastore{db: db}
	})

	return factory, nil
}

//通用操作

func queryByCondition(db *gorm.DB, model interface{}, whereOrder []model.WhereOrder) *gorm.DB {
	tx := db.Model(model)
	if len(whereOrder) > 0 {
		for _, wo := range whereOrder {
			if wo.Order != "" {
				tx = tx.Order(wo.Order)
			}
			if wo.Where != "" {
				tx = tx.Where(wo.Where, wo.Value...)
			}
		}
	}
	return tx
}

func batchDelete(db *gorm.DB, model interface{}, ids []uint64) error {
	return db.Where("id in ?", ids).Delete(model).Error
}
