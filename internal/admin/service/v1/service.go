package v1

import "go-web/internal/admin/store"

type Service interface {
	SysUser() SysUserSrv
	SysRole() SysRoleSrv
	SysMenu() SysMenuSrv
}

type service struct {
	factory store.Factory
}

//工厂模式，创建service
func NewService(factory store.Factory) Service {
	return &service{
		factory: factory,
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
