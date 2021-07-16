package main

import "fmt"

func main() {
	//r := [...]int{5: -1}
	//fmt.Println(r)

	var x = [32]byte{5, 6}
	a := 5
	fmt.Println(x, a)
	fmt.Println(&x, &a)

	zero(&x)
	fmt.Println(x)
	y := &x
	fmt.Println(y)
}
func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
