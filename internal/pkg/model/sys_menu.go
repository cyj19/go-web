package model

type SysMenu struct {
	Model
	URL      string `gorm:"column:url;size:72;"  json:"url" form:"url"`
	Method   string `gorm:"column:method;size:32;" json:"method" form:"method"`
	Name     string `gorm:"column:name;size:32;not null;"  json:"name" form:"name"`
	Sequence int    `gorm:"column:sequence;not null;" json:"sequence" form:"sequence"`
	Type     uint8  `gorm:"column:type;type:tinyint(1);not null;" json:"type" form:"type" `                  //菜单类型 1模块2菜单3操作
	Code     string `gorm:"column:code;size:32;not_null;unique_index:uk_menu_code;" json:"code" form:"code"` //菜单代码
	Icon     string `gorm:"column:icon;size:32;" json:"icon" form:"icon"`
	Status   uint8  `gorm:"column:status;type:tinyint(1);not null;"  json:"status" form:"status"`
	Memo     string `gorm:"column:memo;size:64;"  json:"memo" form:"memo"`
	ParentID uint64 `gorm:"column:parent_id;not null;" json:"parent_id" form:"parent_id"`
}

func (m *SysMenu) TableName() string {
	return "sys_menu"
}

type SysMenuPage struct {
	SysMenu
	PageInfo
}
