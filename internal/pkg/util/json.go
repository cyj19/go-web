package util

import (
	"encoding/json"
	"fmt"
)

// struct 转 json
func Struct2Json(obj interface{}) string {
	data, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Struct2Json函数异常：%v \n", err)
	}
	return string(data)
}

// json 转 struct，因为要修改obj，所以obj必须是指针
func Json2Struct(str string, obj interface{}) {
	err := json.Unmarshal([]byte(str), obj)
	if err != nil {
		fmt.Printf("Json2Struct函数异常：%v \n", err)
	}
}

// struct转struct，根据json标签转换，dest必须是指针
func Struct2Struct(sour, dest interface{}) {
	str := Struct2Json(sour)
	Json2Struct(str, dest)
}
