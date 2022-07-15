package main

import (
	"fmt"
	"math"
)

func Going(n int64) float64 {
	fmt.Println(factorial(n))
	fmt.Println(sumFactorial(n))
	//fmt.Println(float64(1) / float64(factorial(n)))

	//fmt.Println((float64(1) / float64(factorial(n))) * float64(sumFactorial(n)))

	res := (float64(1) / float64(factorial(n))) * float64(sumFactorial(n))

	return math.Floor(res*1000000) / 1000000
}

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	if n-1 < 1 {
		return n
	}
	return n * factorial(n-1)
}

func sumFactorial(n int64) int64 {
	var sum, j int64
	j = 1
	for i := j; i <= n; i++ {
		sum += factorial(i)
	}
	return sum
}

//Calculate (1 / n!) * (1! + 2! + 3! + ... + n!) for a given n, where n is an integer greater or equal to 1.
