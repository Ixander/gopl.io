package main

import (
	"fmt"
	"time"
)

//define the download() function
func download(s int, ch chan int) {
	sum := 0
	for i := 0; i <= s; i++ {
		sum += i
	}
	//fmt.Printf("sum: %d\n", sum)
	ch <- sum
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	var s1, s2, s3 int
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	fmt.Scanln(&s3)

	go download(s1, ch1)
	go download(s2, ch2)
	go download(s3, ch3)
	res := 0
	time.Sleep(150 * time.Millisecond)
	//for i := 0; i < 3; i++ {
	for {
		select {
		case x := <-ch1:
			res += x
			//fmt.Println(x)
		case y := <-ch2:
			res += y
			//fmt.Println(y)
		case z := <-ch3:
			res += z
			//fmt.Println(z)
		default:
			fmt.Println(res)
			return
		}
		//output the sum of all results

	}
	//fmt.Println(res)
}
