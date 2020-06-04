package helper

import (
	"strings"
)

func Trim_Word(toke string) string {
	// 去除空格
	toke = strings.Replace(toke, " ", "", -1)
	// 去除换行符
	toke = strings.Replace(toke, "\n", "", -1)
	return toke
}

