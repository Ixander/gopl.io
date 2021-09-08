package main

import "fmt"

func main() {
	x := 5

	fmt.Println(x)
	fmt.Printf("variable x: %v\n", x)
	res := fmt.Sprintf("Sprintf variable x: %v", x)

	fmt.Println(res)

}
