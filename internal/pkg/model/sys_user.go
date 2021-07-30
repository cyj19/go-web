package model

type SysUser struct {
	Model
	Username string ` gorm:"column:username;size:32;not null;comment:'用户名'" json:"username" form:"username" `
	Password string ` gorm:"column:password;size:64;not null;comment:'密码'" json:"password" form:"password" `
	Status   *bool  ` gorm:"column:status;type:tinyint(1);default:1;comment:'用户状态(0：禁用，1：启动，默认1)'" json:"status" form:"status"`
}

//重命名表名
func (u *SysUser) TableName() string {
	return "sys_user"
}

//分页查询参数接收体
type SysUserPage struct {
	SysUser
	PageInfo
}
