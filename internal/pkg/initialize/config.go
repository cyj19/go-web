package initialize

import (
	"bytes"
	"flag"
	"fmt"
	"go-web/internal/pkg/config"
	"log"
	"strings"

	"github.com/spf13/viper"
)

const (
	confileType = "yml"
	configPath  = "../../configs"
)

/*
	初始化配置文件
	参数1: 默认开发配置文件
	参数2: 默认生产配置文件
*/
func Config(developmentConfig string, productionConfig string) (*config.CustomConfBox, *config.Configuration) {
	var box config.CustomConfBox
	var conf config.Configuration
	// 声明命令行标志
	confFlag := flag.String("web_conf", "", "config path")
	modeFlag := flag.String("web_mode", "", "run mode")
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
	readConfig(&box, developmentConfig)
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
		readConfig(&box, configName)
	}

	// 配置转为结构体
	if err := box.ViperIns.Unmarshal(&conf); err != nil {
		panic(fmt.Sprintf("配置转结构体失败：%v , 配置文件所在目录为：%s", err, box.ConfEnv))
	}

	log.Println("初始化配置文件完成...")

	return &box, &conf

}

func readConfig(box *config.CustomConfBox, filename string) {
	box.ViperIns.SetConfigType(confileType)
	config, err := box.Find(filename)
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败：%v , 配置文件所在目录为：%s", err, box.ConfEnv))
	}
	// 初始化配置
	err = box.ViperIns.ReadConfig(bytes.NewReader(config))
	if err != nil {
		panic(fmt.Sprintf("初始化配置失败：%v , 配置文件所在目录为：%s", err, box.ConfEnv))
	}
}
