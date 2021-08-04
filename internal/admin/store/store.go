package store

// Factory defines the my-admin storage interface
type Factory interface {
	SysUser() SysUserStore
	SysRole() SysRoleStore
	SysMenu() SysMenuStore
	SysApi() SysApiStore
	// 创建操作可以共用
	Create(value interface{}) error
}
