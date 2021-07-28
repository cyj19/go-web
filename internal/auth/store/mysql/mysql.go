package mysql

import (
	"fmt"

	"go-web/internal/auth/store"
	"go-web/internal/pkg/model"
	"go-web/pkg/db"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

var (
	factory           store.Factory
	dbIns             *gorm.DB
	once, onceFactory sync.Once
)

func (d *database) User() store.UserStore {
	return newUser(d)
}

func GetMySQLFactory() (store.Factory, error) {
	db, err := getMySQLIns()
	if db == nil || err != nil {
		return nil, err
	}

	onceFactory.Do(func() {
		factory = &database{db: db}
	})

	return factory, nil
}

func getMySQLIns() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		option := &db.MySQLOption{
			Host:                  viper.GetString("mysql.host"),
			Username:              viper.GetString("mysql.username"),
			Password:              viper.GetString("mysql.password"),
			Database:              viper.GetString("mysql.database"),
			MaxIdleConnections:    viper.GetInt("mysql.max-idle-connections"),
			MaxOpenConnections:    viper.GetInt("max-open-connections"),
			MaxConnectionLifeTime: viper.GetDuration("max-connection-life-time"),
			LogLevel:              viper.GetInt("log-level"),
		}
		dbIns, err = db.NewMySQL(option)
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get mysql db, error: %w", err)
	}

	return dbIns, nil
}

func MigrateTable() error {
	dbTemp, _ := getMySQLIns()
	if err := dbTemp.AutoMigrate(&model.SysUser{}); err != nil {
		return fmt.Errorf("failed to migrate user table, error: %w", err)
	}

	return nil
}
