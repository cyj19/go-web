package model

type SysCasbin struct {
	Id    uint64 `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"colument:ptype;size:100;index:idx_casbin_unique;comment:'策略类型'" json:"type"`
	V0    string `gorm:"size:100;index:idx_casbin_unique;comment:'角色关键字'" json:"roleKey"`
	V1    string `gorm:"size:100;index:idx_casbin_unique;comment:'资源名称'" json:"uri"`
	V2    string `gorm:"size:100;index:idx_casbin_unique;comment:'请求类型'" json:"method"`
	V3    string `gorm:"size:100;index:idx_casbin_unique;"`
	V4    string `gorm:"size:100;index:idx_casbin_unique;"`
	V5    string `gorm:"size:100;index:idx_casbin_unique;"`
}

func (c *SysCasbin) TableName() string {
	return "sys_casbin"
}

type SysCasbinPage struct {
	SysCasbin
	PageInfo
}

type SysRoleCasbin struct {
	Kyeword string `json:"keyword"` // 按角色，对应casbin的v0
	Path    string `json:"path"`    // 对应casbin的v1
	Method  string `josn:"method"`  // 对应casbin的v2
}
