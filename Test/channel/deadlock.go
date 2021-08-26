package main

import "fmt"

func main() {
	numbers := make(chan int)
	go generateNumbers(1000, numbers)

	for {
		number, ok := <-numbers

		fmt.Println(number, ok)
		if !ok {
			break
		}
	}

}

func generateNumbers(n int, res chan int) {
	for i := 0; i <= n; i++ {
		res <- i * 2
	}
	// close chsnnel
	close(res)
}
