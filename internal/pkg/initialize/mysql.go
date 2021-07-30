package initialize

import (
	"fmt"
	"go-web/internal/pkg/global"
	"go-web/internal/pkg/model"
	"sync"

	"gorm.io/gorm"
)

var (
	dbIns *gorm.DB
	once  sync.Once
)

func MySQL() {
	var err error
	// 单例模式，保证整个生命周期只初始化一次
	once.Do(func() {
		dbIns, err = global.NewMySQL(configuration.Mysql)
	})

	if err != nil {
		panic(fmt.Sprintf("初始化MySQL异常：%v", err))
	}

	autoMigrateTables()

}

//自动迁移表结构
func autoMigrateTables() {
	dbIns.AutoMigrate(
		new(model.SysUser),
		new(model.SysRole),
		new(model.SysMenu),
		new(model.SysCasbin),
	)
}

// 暴露给其他包
func GetMySQLIns() *gorm.DB {
	return dbIns
}
