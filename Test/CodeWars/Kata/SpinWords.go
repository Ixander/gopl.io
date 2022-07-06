package main

import (
	"strings"
)

func SpinWords(str string) string {
	s := strings.Split(str, " ")
	for i := range s {
		if len(s[i]) >= 5 {
			s[i] = reverse(s[i])
		}
	}
	return strings.Join(s, " ")
}

func reverse(str string) string {
	res := ""
	for k := range str {
		res += string(str[len(str)-k-1])
	}
	return res
}
