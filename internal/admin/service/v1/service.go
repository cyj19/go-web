package v1

import (
	"fmt"
	"go-web/internal/admin/store"
	"go-web/internal/pkg/cache"
	"go-web/internal/pkg/model"

	"github.com/casbin/casbin/v2"
)

type Service interface {
	SysUser() SysUserSrv
	SysRole() SysRoleSrv
	SysMenu() SysMenuSrv
	SysApi() SysApiSrv
	SysCasbin() SysCasbinSrv
	Create(value interface{}) error
	GetById(id uint64, model interface{}) error
}

type service struct {
	factory  store.Factory
	enforcer *casbin.Enforcer
}

const defaultSize = 10

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

// model必须是指针
func (s *service) GetById(id uint64, value interface{}) error {
	// 从缓存查询
	var tableName string
	switch v := value.(type) {
	case *model.SysUser:
		tableName = v.TableName()
	case *model.SysRole:
		tableName = v.TableName()
	case *model.SysMenu:
		tableName = v.TableName()
	case *model.SysApi:
		tableName = v.TableName()
	default:
		tableName = ""
	}

	key := fmt.Sprintf("%s:id:%d", tableName, id)
	err := cache.Get(key, value)
	if err != nil {
		err = s.factory.GetById(id, value)
		// 写入缓存
		cache.Set(key, value)
		return err
	}
	return nil
}
