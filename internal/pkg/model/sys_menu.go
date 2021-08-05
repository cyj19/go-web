package model

type SysMenu struct {
	Model
	Name      string    `gorm:"comment:'菜单英文名称'" json:"name"`
	Title     string    `gorm:"comment:'菜单标题(中文名称)'" json:"title"`
	Icon      string    `gorm:"comment:'菜单图标'" json:"icon"`
	Path      string    `gorm:"comment:'菜单前端访问路径'" json:"path"`
	Redirect  string    `gorm:"comment:'重定向路径'" josn:"redirect"`
	Component string    `gorm:"comment:'前端组件路径'" json:"component"`
	Sort      *uint     `gomr:"type:int unsigned;comment:'菜单顺序(同级比较越小越前)'" josn:"sort"`            // 定义为指针类型可以避免默认值为0的情况
	Status    *bool     `gorm:"type:tinyint(1);default:1;comment:'菜单状态(0：禁用，1：启动)'" json:"status"` // 定义为指针类型可以避免默认值为false的情况
	ParentId  uint64    `gorm:"column:parent_id;not null;" json:"parent_id" form:"parent_id"`
	Children  []SysMenu `gorm:"-" json:"children"`
	Roles     []SysRole `gorm:"many2many:sys_role_menu_relation;" json:"roles"` // 角色菜单多对多关系表
}

func (m *SysMenu) TableName() string {
	return "sys_menu"
}

type SysMenuPage struct {
	SysMenu
	PageInfo
}
