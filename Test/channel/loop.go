package main

import (
	"fmt"
)

func main() {
	number := make(chan int)

	go func() {
		number <- 42
	}()

	//time.Sleep(time.Millisecond * 100)

	select {
	case n := <-number:
		fmt.Println(n)
	default:

		fmt.Println("complete nothing")
	}
}
