package global

import (
	"fmt"
	"time"

	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// viper内置了mapstructure, yml文件用"-"区分单词, 转为驼峰方便
type Configuration struct {
	Server *ServerConfiguration `mapstructure:"server" json:"server"`
	Mysql  *MysqlConfiguration  `mapstructure:"mysql" json:"mysql"`
	Redis  *RedisConfiguration  `mapstructure:"redis" json:"redis"`
	Casbin *CasbinConfiguration `mapstructure:"casbin" json:"casbin"`
	Jwt    *JWTConfiguration    `mapstructure:"jwt" json:"jwt"`
	Log    *LogConfiguration    `mapstructure:"log" json:"log"`
}

type ServerConfiguration struct {
	Port       int    `mapstructure:"port" json:"port"`
	Name       string `mapstructure:"name" json:"name"`
	UrlPrefix  string `mapstructure:"url-prefix" json:"urlPrefix"`
	ApiVersion string `mapstructure:"api-version" json:"apiVersion"`
	InitData   bool   `mapstructure:"init-data" json:"initData"`
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

	// gorm 默认会在事务里执行写入操作（创建、更新、删除）
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

type JWTConfiguration struct {
	Realm      string `mapstructure:"realm" json:"realm" `
	Key        string `mapstructure:"key" json:"key"`
	Timeout    int    `mapstructure:"timeout" json:"timeout"`
	MaxRefresh int    `mapstructure:"max-refresh" json:"maxRefresh"`
}

type LogConfiguration struct {
	Path       string        `mapstructure:"path" json:"path"`
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	MaxSize    int           `mapstructure:"max-size" json:"maxSize"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}
