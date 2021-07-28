package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

//定义为常用的yyyy-MM-dd hh:mm:ss
type LocalTime struct {
	time.Time
}

//实现driver中的Valuer接口，将自定义类型值转换为驱动支持的Value类型值
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

//实现sql包中的Scanner接口，会被Rows或Row的Scan方法调用
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

//重写time.Time的MarshalJSON方法
func (t LocalTime) MarshalJSON() ([]byte, error) {
	timeStr := t.Format("2006-01-02 15:04:05")
	//零值判断
	if t.IsZero() {
		timeStr = ""
	}
	formatted := fmt.Sprintf("\"%s\"", timeStr)
	return []byte(formatted), nil
}

func (t LocalTime) UnmarshalJSON(data []byte) error {
	//去除json格式中的双引号
	s := strings.Trim(string(data), "\"")
	//空值判断，因为在MarshalJSON中时间是零值，json为空字符串
	if s == "null" || strings.TrimSpace(s) == "" {
		t = LocalTime{Time: time.Time{}}
		return nil
	}
	t = LocalTime{Time: setLocTime(s)}
	return nil
}

func setLocTime(value string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
	return t
}
