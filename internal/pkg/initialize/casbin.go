package initialize

import (
	"go-web/internal/pkg/global"
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
		global.Log.Error(ctx, "初始化Casbin失败：%s", err)
	}

	global.Log.Info(ctx, "初始化Casbin完成...")
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
		modelPath := global.Conf.Casbin.ModelPath
		config, err := box.Find(modelPath)
		if err != nil {
			global.Log.Error(ctx, "读取策略文件失败：%v", err)
			return
		}
		// 创建模型
		casbinModel := model.NewModel()
		err = casbinModel.LoadModelFromText(string(config))
		if err != nil {
			global.Log.Error(ctx, "加载casbin模型失败：%v", err)
			return
		}
		enforcer, err = casbin.NewEnforcer(casbinModel, a)
		if err != nil {
			return
		}

		// 加载策略
		err = enforcer.LoadPolicy()
		if err != nil {
			global.Log.Error(ctx, "加载策略失败：%v", err)
			return
		}

	})
	return err
}

// 暴露给其他包
func GetEnforcerIns() *casbin.Enforcer {
	return enforcer
}
