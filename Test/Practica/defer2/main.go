package main

import "fmt"

var Global = 5

func main() {
	fmt.Println(Global)
	changeGlobal()
	fmt.Println(Global)
}

func changeGlobal() {
	defer func(checkout int) {
		Global = checkout
	}(Global)
	Global = 10
	fmt.Println(Global)
}
