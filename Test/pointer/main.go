package main

import "fmt"

func main() {
	fmt.Println(p)
	//fmt.Println(f())
}

var p = f()

func f() *int {
	v := 1
	return &v
}
