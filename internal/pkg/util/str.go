package util

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	string分割
*/
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

/*
	string转uint64
*/
func Str2Uint64(str string) uint64 {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

/*
	uint64转string
*/
func Uint642Str(num uint64) string {
	return fmt.Sprintf("%d", num)
}

/*
	string切片转uint64切片
*/
func ConverSliceToUint64(strs []string) ([]uint64, error) {
	arr := make([]uint64, 0)
	for _, value := range strs {
		temp, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		arr = append(arr, uint64(temp))
	}
	return arr, nil
}
