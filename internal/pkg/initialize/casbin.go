package initialize

import (
	"fmt"
	"log"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	enforcer     *casbin.Enforcer
	onceEnforcer sync.Once
)

// 初始化casbin
func Casbin() {
	err := mysqlCasbin()
	if err != nil {
		panic(fmt.Sprintf("初始化Casbin失败：%s", err))
	}
	log.Println("初始化Casbin完成")
}

func mysqlCasbin() error {
	// 初始化数据库适配器
	var err error
	onceEnforcer.Do(func() {
		//casbin内部生成的表为casbin_rule，修改表名为sys_casbin，第二个参数为数据库表前缀
		a, err := gormadapter.NewAdapterByDBUseTableName(dbIns, "", "sys_casbin")
		if err != nil {
			return
		}
		// 读取策略文件
		modelPath := configuration.Casbin.ModelPath
		config, err := box.Find(modelPath)
		if err != nil {
			return
		}
		// 创建模型
		casbinModel := model.NewModel()
		err = casbinModel.LoadModelFromText(string(config))
		if err != nil {
			return
		}
		enforcer, err = casbin.NewEnforcer(casbinModel, a)
		if err != nil {
			return
		}

		// 加载策略
		err = enforcer.LoadPolicy()
		if err != nil {
			return
		}

	})
	return err
}

// 暴露给其他包
func GetEnforcerIns() *casbin.Enforcer {
	return enforcer
}
