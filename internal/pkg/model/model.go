package model

import (
	"go-web/pkg/model"

	"github.com/dgrijalva/jwt-go"
)

// gorm提供的Model没有json tag，因此进行自定义
type Model struct {
	ID        uint64           `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`
	CreatedAt model.LocalTime  `gorm:"column:created_at;type:datetime;not null;" json:"created_at" form:"created_at"`
	UpdatedAt model.LocalTime  `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at"  form:"updated_at"`
	DeletedAt *model.DeletedAt `gorm:"column:deleted_at;type:datetime;" json:"deleted_at" `
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
