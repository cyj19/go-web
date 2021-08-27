package global

import (
	"io/ioutil"

	"github.com/spf13/viper"
)

const (
	MsecLocalTimeFormat = "2006-01-02 15:04:05.000"
)

var (
	Conf Configuration
	Log  *GormZapLogger
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
