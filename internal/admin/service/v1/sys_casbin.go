package v1

import (
	"go-web/internal/pkg/model"

	"github.com/casbin/casbin/v2"
)

type SysCasbinSrv interface {
	GetRoleCasbins(roleCasbin model.SysRoleCasbin) []model.SysRoleCasbin
	BatchDeleteRoleCasbins(roleCasbins []model.SysRoleCasbin) (bool, error)
	CreateRoleCasbin(roleCasbin model.SysRoleCasbin) (bool, error)
	BatchCreateRoleCasbins(roleCasbins []model.SysRoleCasbin) (bool, error)
}

type casbinService struct {
	enforcer *casbin.Enforcer
}

func newCasbinService(s *service) SysCasbinSrv {
	return &casbinService{enforcer: s.enforcer}
}

// 按角色即默认p_type=p，获取符合条件的casbin规则
func (c *casbinService) GetRoleCasbins(roleCasbin model.SysRoleCasbin) []model.SysRoleCasbin {
	rules := c.enforcer.GetFilteredGroupingPolicy(0, roleCasbin.Kyeword, roleCasbin.Path, roleCasbin.Method)
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

// 按角色，删除符合条件的casbin规则
func (c *casbinService) BatchDeleteRoleCasbins(roleCasbins []model.SysRoleCasbin) (bool, error) {
	rules := make([][]string, 0)
	for _, v := range roleCasbins {
		rule := []string{v.Kyeword, v.Path, v.Method}
		rules = append(rules, rule)
	}
	return c.enforcer.RemovePolicies(rules)
}

func (c *casbinService) CreateRoleCasbin(roleCasbin model.SysRoleCasbin) (bool, error) {
	return c.enforcer.AddPolicy(roleCasbin.Kyeword, roleCasbin.Path, roleCasbin.Method)
}

// 按角色, 批量创建casbin规则
func (c *casbinService) BatchCreateRoleCasbins(roleCasbins []model.SysRoleCasbin) (bool, error) {
	rules := make([][]string, 0)
	for _, v := range roleCasbins {
		rule := []string{v.Kyeword, v.Path, v.Method}
		rules = append(rules, rule)
	}
	return c.enforcer.AddPolicies(rules)
}
