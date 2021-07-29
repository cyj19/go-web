package initialize

import (
	"fmt"
	"go-web/internal/pkg/model"
	"go-web/pkg/db"
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
		option := &db.MySQLOption{
			Host:                  box.ViperIns.GetString("mysql.host"),
			Username:              box.ViperIns.GetString("mysql.username"),
			Password:              box.ViperIns.GetString("mysql.password"),
			Database:              box.ViperIns.GetString("mysql.database"),
			MaxIdleConnections:    box.ViperIns.GetInt("mysql.max-idle-connections"),
			MaxOpenConnections:    box.ViperIns.GetInt("max-open-connections"),
			MaxConnectionLifeTime: box.ViperIns.GetDuration("max-connection-life-time"),
			LogLevel:              box.ViperIns.GetInt("log-level"),
		}
		dbIns, err = db.NewMySQL(option)
	})

	if err != nil {
		panic(fmt.Sprintf("初始化MySQL异常：%v", err))
	}

	autoMigrateTables()

}

//自动迁移表结构
func autoMigrateTables() {
	dbIns.AutoMigrate(new(model.SysUser))
}

// 暴露给其他包
func GetMySQLIns() *gorm.DB {
	return dbIns
}
