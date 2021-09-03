package initialize

import (
	"fmt"

	"github.com/vagaryer/go-web/internal/pkg/config"
	"github.com/vagaryer/go-web/internal/pkg/db"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// model为表结构
func MySQL(opt *config.MysqlConfiguration, log gormlogger.Interface, models ...interface{}) *gorm.DB {
	dbIns, err := db.NewMySQL(opt, log)

	if err != nil {
		panic(fmt.Sprintf("初始化MySQL异常：%v", err))
	}

	err = autoMigrateTables(dbIns, models...)
	if err != nil {
		panic(fmt.Sprintf("初始化MySQL异常：%v", err))
	}

	return dbIns
}

//自动迁移表结构
func autoMigrateTables(dbIns *gorm.DB, models ...interface{}) error {
	return dbIns.AutoMigrate(models...)
}
