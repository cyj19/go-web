package model

type SysRole struct {
	Model
	Name   string    `gorm:"column:name;size:32;not null;comment:'角色名称'" json:"name" `
	NameZh string    `gorm:"column:name_zh;size:32;comment:'角色中文名称'"  json:"nameZh" `
	Status *bool     `gorm:"column:status;type:tinyint(1);default:1;comment:'角色状态:(0：禁用，1：启用)';" json:"status"`
	Menus  []SysMenu `gorm:"many2many:sys_role_menu_relation;" json:"menus"` // 角色菜单多对多关系表
	Users  []SysUser `gorm:"many2many:sys_user_role_relation;" json:"users"` // 用户角色多对多关系表
}

//重命名表名
func (r *SysRole) TableName() string {
	return "sys_role"
}

//分页查询参数接收体
type SysRolePage struct {
	SysRole
	PageInfo
}
