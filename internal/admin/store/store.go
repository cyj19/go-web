package store

// Factory defines the my-admin storage interface
type Factory interface {
	SysUser() SysUserStore
	SysRole() SysRoleStore
	SysMenu() SysMenuStore
	SysApi() SysApiStore
	// 以下是公用操作
	Create(value interface{}) error
	GetById(id uint64, model interface{}) error
}
