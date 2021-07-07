package main

import "fmt"

func main() {
	p := new(int) // повертає адрес *int
	*p = 2

	// а - значення, сама змінна то &а -
	a := 5
	fmt.Printf("змінна а: %v, адрес змінної &а: %v\n", a, &a)

	pointer := &a
	fmt.Printf("вказівник pointer: %v, значення на яке вказує *pointer: %v\n", pointer, *pointer)

	fmt.Println(p)
	fmt.Println(*p)
}
