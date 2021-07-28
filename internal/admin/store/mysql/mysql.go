package mysql

import (
	"fmt"
	"sync"

	"go-web/internal/admin/store"
	"go-web/internal/pkg/model"
	"go-web/pkg/db"

	"github.com/spf13/viper"
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

//不能放到pkg包中
var (
	factory           store.Factory
	dbIns             *gorm.DB
	once, onceFactory sync.Once
)

//单例 获取factory
func GetMySQLFactory() (store.Factory, error) {
	dbTemp, err := getMySQLIns()
	if dbTemp == nil || err != nil {
		return nil, err
	}
	onceFactory.Do(func() {
		factory = &datastore{db: dbTemp}
	})

	return factory, nil
}

//单例 获取mysql db
func getMySQLIns() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		option := &db.MySQLOption{
			Host:                  viper.GetString("mysql.host"),
			Username:              viper.GetString("mysql.username"),
			Password:              viper.GetString("mysql.password"),
			Database:              viper.GetString("mysql.database"),
			MaxIdleConnections:    viper.GetInt("mysql.max-idle-connections"),
			MaxOpenConnections:    viper.GetInt("mysql.max-open-connections"),
			MaxConnectionLifeTime: viper.GetDuration("mysql.max-connection-life-time"),
			LogLevel:              viper.GetInt("mysql.log-level"),
		}

		dbIns, err = db.NewMySQL(option)

	})

	if err != nil {
		return nil, fmt.Errorf("failed to get mysql db, error: %w", err)
	}

	return dbIns, nil
}

//初始化表
func MigrateTable() error {
	dbTemp, _ := getMySQLIns()
	if err := dbTemp.AutoMigrate(&model.SysUser{}); err != nil {
		return fmt.Errorf("failed to migrate user table, error: %w", err)
	}
	if err := dbTemp.AutoMigrate(&model.SysRole{}); err != nil {
		return fmt.Errorf("failed to migrate role table, error: %w", err)
	}
	if err := dbTemp.AutoMigrate(&model.SysMenu{}); err != nil {
		return fmt.Errorf("failed to migrate menu table, error: %w", err)
	}
	return nil
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

func delete(db *gorm.DB, id uint64, model interface{}) error {
	return db.Where("id = ?", id).Delete(model).Error
}

func deleteBatch(db *gorm.DB, ids []uint64, model interface{}) error {
	return db.Where("id in (?)", ids).Delete(model).Error
}
