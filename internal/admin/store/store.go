package store

import "go-web/internal/pkg/model"

// Factory defines the my-admin storage interface
type Factory interface {
	SysUser() SysUserStore
	SysRole() SysRoleStore
	SysMenu() SysMenuStore
	SysApi() SysApiStore
	// 以下是公用操作
	Create(value interface{}) error
	GetById(id uint64, value interface{}) error
	GetList(value interface{}, result interface{}, whereOrders ...model.WhereOrder) error
	GetPage(pageIndex int, pageSize int, value interface{}, result interface{}, whereOrder ...model.WhereOrder) (int64, error)
}
