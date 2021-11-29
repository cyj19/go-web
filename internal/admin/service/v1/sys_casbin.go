package v1

import (
	"context"

	"github.com/cyj19/go-web/internal/admin/global"
	"github.com/cyj19/go-web/internal/admin/store"
	"github.com/cyj19/go-web/internal/pkg/model"
)

type SysCasbinSrv interface {
	GetRoleCasbins(ctx context.Context, roleCasbin model.SysRoleCasbin) []model.SysRoleCasbin
	BatchDeleteRoleCasbins(ctx context.Context, roleCasbins []model.SysRoleCasbin) (bool, error)
	CreateRoleCasbin(ctx context.Context, roleCasbin model.SysRoleCasbin) (bool, error)
	BatchCreateRoleCasbins(ctx context.Context, roleCasbins []model.SysRoleCasbin) (bool, error)
}

type casbinService struct {
	factory store.Factory
}

func newCasbinService(s *service) SysCasbinSrv {
	return &casbinService{factory: s.factory}
}

var _ SysCasbinSrv = (*casbinService)(nil)

// GetRoleCasbins 按角色即默认p_type=p，获取符合条件的casbin规则
func (c *casbinService) GetRoleCasbins(ctx context.Context, roleCasbin model.SysRoleCasbin) []model.SysRoleCasbin {
	rules := global.Enforcer.GetFilteredGroupingPolicy(0, roleCasbin.Kyeword, roleCasbin.Path, roleCasbin.Method)
	cs := make([]model.SysRoleCasbin, 0)
	for _, rule := range rules {
		cs = append(cs, model.SysRoleCasbin{
			Kyeword: rule[0],
			Path:    rule[1],
			Method:  rule[2],
		})
	}

	return cs
}

// BatchDeleteRoleCasbins 按角色，删除符合条件的casbin规则
func (c *casbinService) BatchDeleteRoleCasbins(ctx context.Context, roleCasbins []model.SysRoleCasbin) (bool, error) {
	rules := make([][]string, 0)
	for _, v := range roleCasbins {
		rule := []string{v.Kyeword, v.Path, v.Method}
		rules = append(rules, rule)
	}
	return global.Enforcer.RemovePolicies(rules)
}

func (c *casbinService) CreateRoleCasbin(ctx context.Context, roleCasbin model.SysRoleCasbin) (bool, error) {
	return global.Enforcer.AddPolicy(roleCasbin.Kyeword, roleCasbin.Path, roleCasbin.Method)
}

// BatchCreateRoleCasbins 按角色, 批量创建casbin规则
func (c *casbinService) BatchCreateRoleCasbins(ctx context.Context, roleCasbins []model.SysRoleCasbin) (bool, error) {
	rules := make([][]string, 0)
	for _, v := range roleCasbins {
		rule := []string{v.Kyeword, v.Path, v.Method}
		rules = append(rules, rule)
	}
	return global.Enforcer.AddPolicies(rules)
}
