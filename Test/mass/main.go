package main

import "fmt"

func main() {
	//r := [...]int{5: -1}
	//fmt.Println(r)

	slice := make([]int, 4)

	fmt.Println(slice)

	var x = [32]byte{5, 6}
	a := 5
	fmt.Println(x, a)
	fmt.Println(&x, &a)

	zero(&x)
	fmt.Println(x)
	y := &x
	fmt.Println(y)

	c := []int{1, 2, 3}
	b := []int{4, 5, 6}
	fmt.Println(c)
	fmt.Println(b)
	copy(c, b) // копіює b в c

	fmt.Println(copy(c, b))
	fmt.Println(c)
	fmt.Println(b)

	b = append(b, 1, 2)

	b = append(b, b...)

	fmt.Println(b)
}
func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
