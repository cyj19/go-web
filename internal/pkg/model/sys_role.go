package model

type SysRole struct {
	Model
	Name   string `gorm:"column:name;size:32;not null;comment:'角色名称'" json:"name" form:"name"`
	NameZh string `gorm:"column:name_zh;size:32;comment:'角色中文名称'"  json:"name_zh" form:"name_zh"`
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
