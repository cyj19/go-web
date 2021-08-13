package model

type SysApi struct {
	Model
	Method   string   `gorm:"not null;comment:'请求方式'" json:"method"`
	Path     string   `gorm:"not null;comment:'访问路径'" json:"path"`
	Category string   `gorm:"comment:'所属类别'" json:"category"`
	Desc     string   `gorm:"comment:'说明'" json:"desc"`
	Creator  string   `gorm:"comment:'创建人'" json:"creator"`
	Roles    []uint64 `gorm:"-" json:"roles"`
}

func (a *SysApi) TableName() string {
	return "sys_api"
}

type SysApiPage struct {
	SysApi
	PageInfo
}
