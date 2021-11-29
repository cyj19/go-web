package v1

import (
	"github.com/cyj19/go-web/internal/admin/global"
	"github.com/cyj19/go-web/internal/admin/store"
	"github.com/cyj19/go-web/internal/pkg/cache"
)

type Service interface {
	SysUser() SysUserSrv
	SysRole() SysRoleSrv
	SysMenu() SysMenuSrv
	SysApi() SysApiSrv
	SysCasbin() SysCasbinSrv
}

type service struct {
	factory store.Factory
}

const defaultSize = 10

// NewService 创建service
func NewService(factory store.Factory) Service {
	return &service{
		factory: factory,
	}
}

// 用于触发编译期的接口的合理性检查机制
var _ Service = (*service)(nil)

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

func cleanCache(pattern string) error {
	keys := cache.Keys(global.RedisIns, pattern)
	return cache.Del(global.RedisIns, keys...)
}
