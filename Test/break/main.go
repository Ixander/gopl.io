package main

import "fmt"

func main() {
	a := 0
	for {
		a++
		if a == 15 {
			break
		}

		if a == 10 {
			continue
		}
		fmt.Println(a)
		//a++
	}
}
