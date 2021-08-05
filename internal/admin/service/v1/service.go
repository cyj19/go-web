package v1

import (
	"go-web/internal/admin/store"

	"github.com/casbin/casbin/v2"
)

type Service interface {
	SysUser() SysUserSrv
	SysRole() SysRoleSrv
	SysMenu() SysMenuSrv
	SysApi() SysApiSrv
	SysCasbin() SysCasbinSrv
	Create(value interface{}) error
}

type service struct {
	factory  store.Factory
	enforcer *casbin.Enforcer
}

//工厂模式，创建service
func NewService(factory store.Factory, enforcer *casbin.Enforcer) Service {
	return &service{
		factory:  factory,
		enforcer: enforcer,
	}
}

func (s *service) SysUser() SysUserSrv {
	//创建userService
	return newSysUser(s)
}

func (s *service) SysRole() SysRoleSrv {
	return newSysRole(s)
}

func (s *service) SysMenu() SysMenuSrv {
	return newSysMenu(s)
}

func (s *service) SysApi() SysApiSrv {
	return newSysApi(s)
}

func (s *service) SysCasbin() SysCasbinSrv {
	return newCasbinService(s)
}

func (s *service) Create(value interface{}) error {
	return s.factory.Create(value)
}
