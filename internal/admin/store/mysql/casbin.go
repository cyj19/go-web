package mysql

import (
	"fmt"
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	enforcer     *casbin.Enforcer
	onceEnforcer sync.Once
)

//单例
func GetEnforcerIns() (*casbin.Enforcer, error) {
	var err error
	onceEnforcer.Do(func() {
		db, _ := getMySQLIns()
		a, _ := gormadapter.NewAdapterByDB(db)
		enforcer, err = casbin.NewEnforcer("rbac_model.conf", a)
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get enforcer, error: %w", err)
	}
	return enforcer, nil
}
