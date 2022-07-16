package main

import (
	"math"
	"math/big"
)

func Going(n int) float64 {
	res := big.NewFloat(1.0)
	fact := factorial(n)
	sum := sumFactorial(n)
	res = res.Quo(sum, fact)
	resFloat, _ := res.Float64()
	return math.Floor(resFloat*1e6) / 1e6
}

func factorial(n int) *big.Float {
	factVal := big.NewFloat(1.0)
	for i := 1; i <= n; i++ {
		factVal = factVal.Mul(factVal, big.NewFloat(float64(i)))
	}
	return factVal
}

func sumFactorial(n int) *big.Float {
	sum := big.NewFloat(0.0)
	for i := 1; i <= n; i++ {
		sum = sum.Add(sum, factorial(i))
	}
	return sum
}

//Calculate (1 / n!) * (1! + 2! + 3! + ... + n!) for a given n, where n is an integer greater or equal to 1.
