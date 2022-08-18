package main

import (
	"fmt"
)

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	fmt.Println(RemoveDuplicates(input))
}

//RemoveDuplicates
//Напишите функцию, которая убирает дубликаты, сохраняя порядок слайса:
func RemoveDuplicates(input []string) []string {
	//m := make(map[string]string)
	//res := input
	for i, a := range input {
		for j := i + 1; j < len(input); j++ {
			if a == input[j] {
				//delete an element from a slice
				copy(input[j:], input[j+1:]) // Shift a[i+1:] left one index.
				input[len(input)-1] = ""     // Erase last element (write zero value).
				input = input[:len(input)-1]
				//fmt.Println(a, j)
			}
		}
	}
	return input
}
