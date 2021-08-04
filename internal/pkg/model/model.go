package model

import (
	"go-web/pkg/model"

	"github.com/dgrijalva/jwt-go"
)

// gorm提供的Model没有json tag，因此进行自定义
type Model struct {
	Id        uint64           `gorm:"primaryKey;comment:'自增编号'" json:"id"`
	CreatedAt model.LocalTime  `gorm:"comment:'创建时间'" json:"createdAt"`
	UpdatedAt model.LocalTime  `gorm:"comment:'更新时间'" json:"updatedAt"`
	DeletedAt *model.DeletedAt `gorm:"index:idx_deleted_at;comment:'删除时间(软删除)'" json:"deletedAt"`
}

type Claims struct {
	UserId uint64
	jwt.StandardClaims
}

//sql的条件
type WhereOrder struct {
	Order string
	Where string
	Value []interface{}
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateDelete struct {
	Id     uint64   `json:"id"`     // 需更新记录的id
	Create []uint64 `json:"create"` // 需删除的关联id (角色id 或 菜单id 或 接口id)
	Delete []uint64 `json:"delete"` // 需增加的关联id (角色id 或 菜单id 或 接口id)
}

type IdParam struct {
	Ids string `json:"ids" form:"ids"`
}
