package main

import (
	"fmt"
	"sort"
)

func FindOdd(seq []int) int {

	res := 0
	for _, x := range seq {
		//res ^= x
		//fmt.Println(res)
		res = res ^ x
	}
	fmt.Println(res)

	sort.Ints(seq)

	for i := range seq {
		if i == len(seq)-1 {
			break
		}
		if (seq[i] != seq[i+1]) && ((i+1)%2 != 0) {
			return seq[i]
		}
	}
	return 0
	//return reduce(seq)[0]
}

func reduce(mass []int) []int {
	first := mass[0]
	count := 0
	for _, v := range mass {
		if v == first {
			count++
		}
	}
	if count%2 == 0 {
		mass = removeElement(first, mass)
	} else {
		return mass
	}
	return reduce(mass)
}

func removeElement(elem int, slice []int) []int {
	for i, v := range slice {
		if v == elem {

			copy(slice[i:], slice[i+1:])
			slice[len(slice)-1] = 0
			slice = slice[:len(slice)-1]
			break
		} else {
			return slice
		}
	}
	return removeElement(elem, slice)
}
