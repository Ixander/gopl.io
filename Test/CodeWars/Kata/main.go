package main

import (
	"fmt"
	"strings"
)

type MyString string

func main() {
	//var s MyString
	//s = "ABC"
	//fmt.Println(s.IsUpperCase())
	//fmt.Println(Multiple3And5(12555))
	//fmt.Println(SpinWords("This is a Sparta"))
	//fmt.Println(ReverseWords("double  spaces"))
	//massF := []int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5, 1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5, -1}
	//fmt.Println(FindOdd(massF))
	//fmt.Println(DigitalRoot(2569458888455488))
	//fmt.Println(ArrayDiff([]int{1, 2, 2, 2, 2, 2, 2, 2}, []int{2}))
	fmt.Println("---------------------------------")
	fmt.Println(Going(10110))

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
