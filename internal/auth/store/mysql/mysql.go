package mysql

import (
	"fmt"

	"go-web/internal/auth/initialize"
	"go-web/internal/auth/store"
	"sync"

	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

var (
	factory           store.Factory
	once, onceFactory sync.Once
)

func (d *database) User() store.UserStore {
	return newUser(d)
}

func GetMySQLFactory() (store.Factory, error) {
	db := initialize.GetMySQLIns()
	if db == nil {
		return nil, fmt.Errorf("MySQL实例对象为空")
	}

	onceFactory.Do(func() {
		factory = &database{db: db}
	})

	return factory, nil
}
