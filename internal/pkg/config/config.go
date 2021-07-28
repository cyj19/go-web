package config

import "github.com/spf13/viper"

// 初始化配置
func Init(path, name, fileType string) error {
	viper.SetConfigName(name)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(path)
	return viper.ReadInConfig()

}

func GetServerPort() int {
	return viper.GetInt("server.port")
}
