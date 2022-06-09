package main

import "fmt"

func main() {
	var (
		mass [5]int
		x    int
		min  int
		max  int
		sum  int
	)

	for i := 0; i < 5; i++ {
		fmt.Scanln(&x)
		mass[i] = x
	}
	min = mass[0]
	max = mass[0]
	for i := 0; i < len(mass); i++ {

		if min > mass[i] {
			min = mass[i]
		}
		if max < mass[i] {
			max = mass[i]
		}
		sum = sum + mass[i]
	}
	fmt.Printf("minimum: %d, maximum: %d", sum-max, sum-min)
}
