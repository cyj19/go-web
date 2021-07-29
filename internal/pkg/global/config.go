package global

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// viper内置了mapstructure, yml文件用"-"区分单词, 转为驼峰方便
type Configuration struct {
	Server *ServerConfiguration
	Mysql  *MysqlConfiguration
	Redis  *RedisConfiguration
	Casbin *CasbinConfiguration
}

type ServerConfiguration struct {
	Port       int    `mapstructure:"port" json:"port"`
	Name       string `mapstructure:"name" json:"name"`
	UrlPrefix  string `mapstructure:"url-prefix" json:"urlPrefix"`
	ApiVersion string `mapstructure:"api-version" json:"apiVersion"`
}

type MysqlConfiguration struct {
	Host                  string        `mapstructure:"host" json:"host"`
	Username              string        `mapstructure:"username" json:"username"`
	Password              string        `mapstructure:"password" json:"password"`
	Database              string        `mapstructure:"database" json:"database"`
	MaxIdleConnections    int           `mapstructure:"max-idle-connections" json:"maxIdleConnections"`
	MaxOpenConnections    int           `mapstructure:"max-open-connections" json:"maxOpenConnections"`
	MaxConnectionLifeTime time.Duration `mapstructure:"max-connection-life-time" json:"maxConnectionLifeTime"`
	LogLevel              int           `mapstructure:"log-level" json:"logLevel"`
}

// 根据MysqlConfiguration打开一个数据库连接
func NewMySQL(opt *MysqlConfiguration) (*gorm.DB, error) {
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

type RedisConfiguration struct {
	Addr     string `mapstructure:"addr" json:"addr"`
	Password string `mapstructure:"password" json:"password"`
	Db       int    `mapstructure:"db" json:"db"`
}

type CasbinConfiguration struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath"`
}
