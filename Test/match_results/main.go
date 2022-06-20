package main

import "fmt"

func main() {
	var input string
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	fmt.Scanln(&input)
	results = append(results, input)
	fmt.Println(calculate(results))
}

func calculate(mass []string) int {
	correlation := map[string]int{"w": 3, "d": 1, "l": 0}
	sum := 0
	for _, v := range mass {
		sum += correlation[v]
	}
	return sum
}
