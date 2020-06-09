package helper

import (
	"strings"
)

func Trim_Word(toke string) string {
	return strings.TrimSpace(toke)
}

func Space_Word(toke string) int {
	toke = strings.Replace(toke, "\n", "", -1)
	var s int
	for i := 0; i < len(toke); i++ {
		switch {
			case toke[i] == 32:
				s += 1
			default:
				return s
		}
	}
	return s
}

func Copy_Word(toke string,str string,n int) string {
	if n>0{
		for i := 0; i < n; i++ {
			toke = toke + str
		}
	}else{
		for i := n; i < 0; i++ {
			toke = str + toke
		}
	}
	return toke
}

