package util

import "strconv"

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

/*
	string转uint64
*/

func StrToUint64(str string) uint64 {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}
