package initialize

import (
	"fmt"
	"go-web/internal/pkg/config"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// 初始化casbin
func Casbin(dbIns *gorm.DB, box *config.CustomConfBox, conf *config.Configuration) *casbin.Enforcer {
	enforcer, err := newCasbin(dbIns, box, conf)
	if err != nil {
		panic(fmt.Sprintf("初始化casbin失败：%s", err))
	}
	return enforcer
}

func newCasbin(dbIns *gorm.DB, box *config.CustomConfBox, conf *config.Configuration) (*casbin.Enforcer, error) {
	// 初始化数据库适配器
	var err error

	//casbin内部生成的表为casbin_rule，修改表名为sys_casbin，第二个参数为数据库表前缀
	a, err := gormadapter.NewAdapterByDBUseTableName(dbIns, "", "sys_casbin")
	if err != nil {
		return nil, err
	}
	// 读取策略文件
	modelPath := conf.Casbin.ModelPath
	config, err := box.Find(modelPath)
	if err != nil {
		return nil, err
	}
	// 创建模型
	casbinModel := model.NewModel()
	err = casbinModel.LoadModelFromText(string(config))
	if err != nil {
		return nil, err
	}
	enforcer, err := casbin.NewEnforcer(casbinModel, a)
	if err != nil {
		return nil, err
	}

	// 加载策略
	err = enforcer.LoadPolicy()
	if err != nil {
		//global.Log.Error(ctx, "加载策略失败：%v", err)
		return nil, err
	}

	return enforcer, nil
}
