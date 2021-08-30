package store

import (
	"fmt"
	"go-web/internal/pkg/model"
	"sync"

	"gorm.io/gorm"
)

// Factory defines the my-admin storage interface
type Factory interface {
	SysUser() SysUserStore
	SysRole() SysRoleStore
	SysMenu() SysMenuStore
	SysApi() SysApiStore
	// values必须是指针
	Create(values interface{}) error
	// value必须是指针
	BatchDelete(ids []uint64, value interface{}) error
	// values必须是指针
	Update(values interface{}) error
	// value必须是指针
	GetById(id uint64, value interface{}) error
	// value必须是指针，result必须是指针
	GetList(value interface{}, result interface{}, whereOrders ...model.WhereOrder) error
	// value必须是指针，result必须是指针
	GetPage(pageIndex int, pageSize int, value interface{}, result interface{}, whereOrder ...model.WhereOrder) (int64, error)
}

type datastore struct {
	db *gorm.DB
}

//实现Factory接口
func (ds *datastore) SysUser() SysUserStore {
	return newSysUser(ds)
}

func (ds *datastore) SysRole() SysRoleStore {
	return newSysRole(ds)
}

func (ds *datastore) SysMenu() SysMenuStore {
	return newSysMenu(ds)
}

func (ds *datastore) SysApi() SysApiStore {
	return newSysApi(ds)
}

// value必须是指针
func (ds *datastore) Create(values interface{}) error {
	return ds.db.Create(values).Error
}

// value必须是指针
func (ds *datastore) BatchDelete(ids []uint64, value interface{}) error {
	return ds.db.Where("id in ?", ids).Delete(value).Error
}

// Updates使用 struct 更新时，默认情况下，GORM 只会更新非零值的字段
func (ds *datastore) Update(values interface{}) error {
	return ds.db.Updates(values).Error
}

// result要用于绑定数据，必须是指针类型
func (ds *datastore) GetById(id uint64, result interface{}) error {
	db := ds.db
	switch result.(type) {
	case *model.SysUser:
		db = db.Preload("Roles")
	case *model.SysRole:
		db = db.Preload("Menus").Order("sort")
	case *model.SysMenu:
		db = db.Order("parent_id, sort")
	}
	return db.Where("id = ?", id).First(result).Error
}

// value用于设置db.Model，必须是指针；result用于绑定数据，必须是指针
func (ds *datastore) GetList(value interface{}, result interface{}, whereOrders ...model.WhereOrder) error {
	var db *gorm.DB
	switch v := value.(type) {
	case *model.SysUser:
		db = queryByCondition(ds.db, v, whereOrders).Preload("Roles")
	case *model.SysRole:
		db = queryByCondition(ds.db, v, whereOrders).Preload("Menus").Order("sort")
	case *model.SysMenu:
		db = queryByCondition(ds.db, v, whereOrders).Order("parent_id, sort")
	default:
		db = queryByCondition(ds.db, v, whereOrders)
	}

	return db.Find(result).Error

}

// value用于设置db.Model，必须是指针；result用于绑定数据，必须是指针
func (ds *datastore) GetPage(pageIndex int, pageSize int, value interface{}, result interface{}, whereOrders ...model.WhereOrder) (int64, error) {
	var db *gorm.DB
	switch v := value.(type) {
	case *model.SysUser:
		db = queryByCondition(ds.db, v, whereOrders).Preload("Roles")
	case *model.SysRole:
		db = queryByCondition(ds.db, v, whereOrders).Preload("Menus").Order("sort")
	case *model.SysMenu:
		db = queryByCondition(ds.db, v, whereOrders).Order("parent_id, sort")
	default:
		db = queryByCondition(ds.db, v, whereOrders)
	}

	//查询总记录数
	var count int64
	var err error
	err = db.Count(&count).Error
	if err != nil || count == 0 {
		return count, err
	}
	err = db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(result).Error
	return count, err
}

//不能放到pkg包中
var (
	factory     Factory
	onceFactory sync.Once
)

//单例 获取factory
func NewSqlFactory(db *gorm.DB) (Factory, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
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
