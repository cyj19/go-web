package store

import "go-web/internal/pkg/model"

// Factory defines the my-admin storage interface
type Factory interface {
	SysUser() SysUserStore
	SysRole() SysRoleStore
	SysMenu() SysMenuStore
	SysApi() SysApiStore
	// values必须是指针
	Create(values interface{}) error
	// value必须是指针
	BatchDelete(ids []uint64, value interface{}) error
	// values必须是指针
	Update(values interface{}) error
	// value必须是指针
	GetById(id uint64, value interface{}) error
	// value必须是指针，result必须是指针
	GetList(value interface{}, result interface{}, whereOrders ...model.WhereOrder) error
	// value必须是指针，result必须是指针
	GetPage(pageIndex int, pageSize int, value interface{}, result interface{}, whereOrder ...model.WhereOrder) (int64, error)
}
