package main

import "fmt"

func main() {

	v := 1
	incr(&v)
	incr(&v)
	fmt.Println(incr((&v)))
	//fmt.Println(f())
}

var p = f()

func f() *int {
	v := 1
	return &v
}
func incr(p *int) int {
	*p++
	return *p
}
