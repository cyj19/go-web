package config

import (
	"io/ioutil"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

const (
	MsecLocalTimeFormat = "2006-01-02 15:04:05.000"
)

//自定义配置盒子，存放环境配置和对应的viper
type CustomConfBox struct {
	// 配置文件所在文件目录
	ConfEnv string
	// 配置实例
	ViperIns *viper.Viper
}

//查找配置文件
func (c *CustomConfBox) Find(filename string) ([]byte, error) {

	return ioutil.ReadFile(c.ConfEnv + "/" + filename)

}

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
