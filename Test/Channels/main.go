package main

import "fmt"

func f(ch chan int) {

	ch <- 8

}

func main() {

	ch := make(chan int)

	go f(ch)

	fmt.Println(<-ch)
	select {

	case <-ch:

		fmt.Println("a")

	default:

		fmt.Println("b")

	}

}
