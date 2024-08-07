package utils

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

//@author: dzsdbsdxq
//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// CalcPages
// @author: dzsdbsdxq
// @function: CalcPages
// @description: 计算页数
// @param: total，size int64
// @return: int64
func CalcPages(total, size int64) int64 {
	pages := total / size
	if total%size != 0 {
		pages++
	}
	return pages
}

func EmailVerify(email string) bool {
	result, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	if result {
		return true
	} else {
		return false
	}
}
