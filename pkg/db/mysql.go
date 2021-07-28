package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Option defines option for mysql database
type MySQLOption struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
}

// New create a new gorm db instance with the given MySQLOption
func NewMySQL(opt *MySQLOption) (*gorm.DB, error) {
	dns := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opt.Username,
		opt.Password,
		opt.Host,
		opt.Database,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{Logger: logger.Default.LogMode(logger.LogLevel(opt.LogLevel))})
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
