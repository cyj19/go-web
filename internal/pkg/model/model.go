package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//定义为常用的yyyy-MM-dd hh:mm:ss
type LocalTime struct {
	time.Time
}

// 实现driver中的Valuer接口，将自定义类型值转换为驱动支持的Value类型值，ps: 接收器不能是指针类型，否则会报invalid memory address or nil pointer dereference
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 实现sql包中的Scanner接口，会被Rows或Row的Scan方法调用
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 重写time.Time的MarshalJSON方法
func (t LocalTime) MarshalJSON() ([]byte, error) {
	timeStr := t.Format("2006-01-02 15:04:05")
	//零值判断
	if t.IsZero() {
		timeStr = ""
	}
	formatted := fmt.Sprintf("\"%s\"", timeStr)
	return []byte(formatted), nil
}

// 重写time.Time的UnmarshalJSON方法
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	//去除json格式中的双引号
	s := strings.Trim(string(data), "\"")
	//空值判断，因为在MarshalJSON中时间是零值，json为空字符串
	if s == "null" || strings.TrimSpace(s) == "" {
		*t = LocalTime{Time: time.Time{}}
		return nil
	}
	*t = LocalTime{Time: setLocTime(s)}
	return nil
}

func setLocTime(value string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
	if err != nil {
		return time.Now()
	}

	return t
}

// gorm提供的Model没有json tag，因此进行自定义
type Model struct {
	Id        uint64     `gorm:"primaryKey;comment:'自增编号'" json:"id"`
	CreatedAt LocalTime  `gorm:"comment:'创建时间'" json:"createdAt"`
	UpdatedAt LocalTime  `gorm:"comment:'更新时间'" json:"updatedAt"`
	DeletedAt *DeletedAt `gorm:"index:idx_deleted_at;comment:'删除时间(软删除)'" json:"deletedAt"`
}

type Claims struct {
	UserId uint64
	jwt.StandardClaims
}

//sql的条件，可以自由再添加or,limit,offset
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
