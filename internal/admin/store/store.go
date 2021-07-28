package store

// Factory defines the my-admin storage interface
type Factory interface {
	SysUser() SysUserStore
	SysRole() SysRoleStore
	SysMenu() SysMenuStore
}
