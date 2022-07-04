package main

import (
	"fmt"
	"strings"
)

func main() {
	x := 5
	y := 6
	fmt.Println(multiply(x, y))
	fmt.Println(Opposite(x))
	str := "1.AbcD"
	ToAlternatingCase(str)
	mass := []int{1, 2, 3, -8}
	PositiveSum(mass)

	AbbrevName("patrick feeney")
}

func multiply(x, y int) int {
	return x * y
}
func Opposite(value int) int {
	return -value
}

func ToAlternatingCase(str string) string {
	//fmt.Println(str[0])
	res := ""
	for _, value := range str {
		if string(value) == strings.ToUpper(string(value)) {
			res += strings.ToLower(string(value))
		} else {
			res += strings.ToUpper(string(value))
		}
	}
	fmt.Println(res)
	return res
}

func PositiveSum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		if value > 0 {
			sum += value
		}
	}
	fmt.Println(sum)
	return sum
}

func AbbrevName(name string) string {
	//your code here
	s := strings.Split(name, " ")
	return strings.ToUpper(string(s[0][0])) + "." + strings.ToUpper(string(s[1][0]))
}
