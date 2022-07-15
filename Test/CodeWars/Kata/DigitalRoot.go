package main

import (
	"strconv"
)

func DigitalRoot(n int) int {
	t := strconv.Itoa(n)
	sum := 0
	for _, v := range t {
		d, _ := strconv.Atoi(string(v))
		sum += d
	}
	if sum < 10 {
		return sum
	}
	return DigitalRoot(sum)
}
