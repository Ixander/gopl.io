package main

import "strings"

func ReverseWords(str string) string {

	s := strings.Split(str, " ")
	for i := range s {
		s[i] = reverse1(s[i])
	}
	return strings.Join(s, " ")
}

func reverse1(str string) string {
	res := ""
	for k := range str {
		res += string(str[len(str)-k-1])
	}
	return res
}
