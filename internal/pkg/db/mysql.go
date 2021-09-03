package db

import (
	"fmt"

	"github.com/vagaryer/go-web/internal/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 根据MysqlConfiguration打开一个数据库连接
func NewMySQL(opt *config.MysqlConfiguration, log logger.Interface) (*gorm.DB, error) {
	dns := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opt.Username,
		opt.Password,
		opt.Host,
		opt.Database,
		true,
		"Local")

	// gorm 默认会在事务里执行写入操作（创建、更新、删除）
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{Logger: log})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(opt.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(opt.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(opt.MaxConnectionLifeTime)
	return db, nil
}
