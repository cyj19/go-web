package util

import (
	"strconv"
	"strings"
)

//字符串分割
func Str2Uint64Array(str string) []uint64 {
	arr := strings.Split(str, ",")
	resutl := make([]uint64, 0)
	if len(arr) > 0 {
		for _, v := range arr {
			resutl = append(resutl, Str2Uint64(v))
		}
	} else {
		resutl = append(resutl, Str2Uint64(str))
	}

	return resutl
}

// 字符串转uint64
func Str2Uint64(str string) uint64 {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}
