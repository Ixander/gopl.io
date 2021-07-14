package main

import "fmt"

const (
	i = 1 << (10 * iota)
	KiB
	MiB
	GiB
)

func main() {
	fmt.Println(i)
	fmt.Println(KiB)
	fmt.Println(MiB)
	fmt.Println(GiB)
}
