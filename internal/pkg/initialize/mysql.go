package initialize

import (
	"fmt"
	"go-web/internal/pkg/global"
	"sync"

	"gorm.io/gorm"
)

var (
	dbIns *gorm.DB
	once  sync.Once
)

// model为表结构
func MySQL(models ...interface{}) {
	var err error
	// 单例模式，保证整个生命周期只初始化一次
	once.Do(func() {
		dbIns, err = global.NewMySQL(configuration.Mysql)
	})

	if err != nil {
		panic(fmt.Sprintf("初始化MySQL异常：%v", err))
	}

	autoMigrateTables(models...)

}

//自动迁移表结构
func autoMigrateTables(models ...interface{}) {
	dbIns.AutoMigrate(models...)
}

// 暴露给其他包
func GetMySQLIns() *gorm.DB {
	return dbIns
}
