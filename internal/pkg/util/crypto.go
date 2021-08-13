package util

import (
	"crypto/md5"
	"fmt"
)

//md5加密
func EncryptionPsw(psw string) string {
	data := []byte(psw)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
