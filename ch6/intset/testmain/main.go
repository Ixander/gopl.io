package main

import (
	"GoBook/ch6/intset"
	"fmt"
)

func main() {
	var y intset.IntSet

	x := intset.IntSet{[]uint64{1}}
	fmt.Println(x.String())

	x.Add(1)
	x.Add(144)
	x.Add(9)

	fmt.Println(x.String())
	fmt.Println(y)

}
