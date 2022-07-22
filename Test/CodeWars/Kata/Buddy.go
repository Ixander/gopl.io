package main

import (
	"fmt"
	"math"
)

func Buddy(start, limit int) []int {

	for i := start; i <= limit; i++ {
		sum := divisiorSum(i)
		if i < sum && i == divisiorSum(sum) {
			return []int{i, sum}
		}
	}
	return nil
}

func divisiorSum(n int) int {
	res := 0
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if i == n/i {
				res += i
			} else {
				res += i + n/i
			}
		}
	}
	return res
}

func SumPropDiv(num int) int {
	var sum int

	for i := 2; i < num/i; i++ {

		fmt.Println(i, num/i)
		if num%i == 0 {
			sum += i + num/i
		}
	}
	return sum
}

/*function buddy(start,limit) {
const s = (n) => {
let res = 0;
for(let i = 2; i <= Math.sqrt(n); i++)
if(n % i === 0) i === n/i ? res+= i : res += i + n/i
return res
}

for(let i=start; i<= limit; i++) {
let si = s(i)
if(i < si && i === s(si)) return [i,si]
}
return "Nothing";
}*/

/*func primeFactors(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

// return p^i
func power(p, i int) int {
	result := 1
	for j := 0; j < i; j++ {
		result *= p
	}
	return result
}

// formula comes from https://math.stackexchange.com/a/22723
func SumOfProperDivisors(n int) int {
	pfs := primeFactors(n)
	// key: prime
	// value: prime exponents
	m := make(map[int]int)
	for _, prime := range pfs {
		_, ok := m[prime]
		if ok {
			m[prime] += 1
		} else {
			m[prime] = 1
		}
	}

	sumOfAllFactors := 1
	for prime, exponents := range m {
		sumOfAllFactors *= (power(prime, exponents+1) - 1) / (prime - 1)
	}
	return sumOfAllFactors - n - 1
}
*/
/*func divisiorSum(n int) int {
	sum := 0
	for i := n - 1; i > 1; i-- {
		if n%i == 0 {
			sum += i

		}
	}

	return sum
}*/

/*func divisiorSum2(k, n int) int {
	var sum int
	if n == 1 {
		return sum
	}
	if k == n {
		return sum + divisiorSum2(k, n-1)
	}

	if k%n == 0 {
		return sum + divisiorSum2(k, n-1) + n
	}
	return sum + divisiorSum2(k, n-1)
}*/
