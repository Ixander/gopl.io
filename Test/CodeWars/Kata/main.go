package main

import (
	"fmt"
	"strings"
)

type MyString string

func main() {
	var s MyString
	s = "ABC"
	fmt.Println(s.IsUpperCase())
	fmt.Println(Multiple3And5(12555))
	fmt.Println("---------------------------------")
	fmt.Println(SpinWords("This is a Sparta"))

}

func (s MyString) IsUpperCase() bool {
	for _, value := range s {
		if string(value) != strings.ToUpper(string(value)) {
			return false
		}
	}
	return true
}

func Multiple3And5(number int) int {
	res := 0
	for i := 1; i < number; i++ {
		if (i%3) == 0 || (i%5) == 0 {
			res += i
		}
	}
	return res
}
