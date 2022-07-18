package main

import "fmt"

func main() {

	defer fmt.Println("defer1")
	defer fmt.Println("defer2") //LIFO Stack

	inc := increment()

	fmt.Println(inc())
	//inc = increment()
	fmt.Println(inc())
	//fmt.Println(inc())
	//fmt.Println(inc())

	/*fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())*/

	//inc2 := increment()
	//fmt.Println(inc2())
}

func increment() func() int {
	count := 0
	fmt.Println("count: ", count)
	return func() int {
		count++
		fmt.Println("count++: ", count)
		return count
	}
}

func increment2() int {
	count := 0
	count++
	return count
}
