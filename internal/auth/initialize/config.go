package initialize

import (
	"bytes"
	"flag"
	"fmt"
	"go-web/internal/pkg/global"
	"strings"

	"github.com/spf13/viper"
)

const (
	confileType       = "yml"
	configPath        = "../../config"
	developmentConfig = "auth.dev.yml"
	productionConfig  = "auth.prod.yml"
)

var box = new(global.CustomConfBox)

// 初始化配置文件
func Config() {
	// 声明命令行标志
	confFlag := flag.String("auth_web_conf", "", "config path")
	modeFlag := flag.String("auth_web_mode", "", "run mode")
	flag.Parse()
	// 从命令行中读取配置文件目录
	authWebConf := strings.ToLower(*confFlag)
	if authWebConf == "" {
		//使用默认配置
		authWebConf = configPath
	}
	box.ConfEnv = authWebConf

	box.ViperIns = viper.New()
	// 读取默认配置文件
	readConfig(developmentConfig)
	// 把开发配置作为默认配置
	settings := box.ViperIns.AllSettings()
	for key, value := range settings {
		box.ViperIns.SetDefault(key, value)
	}

	// 获取当前环境模式
	env := strings.ToLower(*modeFlag)
	configName := ""
	if env == "prod" {
		configName = productionConfig
	}

	if configName != "" {
		// 重新读取配置文件，修改和默认配置不同的部分
		readConfig(configName)
	}

}

func readConfig(filename string) {
	box.ViperIns.SetConfigType(confileType)
	config, err := box.Find(filename)
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败：%v , 配置文件路径为：%s", err, box.ConfEnv+"/"+filename))
	}
	// 初始化配置
	err = box.ViperIns.ReadConfig(bytes.NewReader(config))
	if err != nil {
		panic(fmt.Sprintf("初始化配置失败：%v , 配置文件路径为：%s", err, box.ConfEnv+"/"+filename))
	}
}

func GetCustomServer() *global.CustomServer {
	return &global.CustomServer{
		Port:       box.ViperIns.GetInt("server.port"),
		Name:       box.ViperIns.GetString("server.name"),
		UrlPrefix:  box.ViperIns.GetString("server.url-prefix"),
		ApiVersion: box.ViperIns.GetString("server.api-version"),
	}
}
