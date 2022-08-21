package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main_start")
	defer metricTime(time.Now())
	time.Sleep(1 * time.Second)

	defer fmt.Println("defer1")
	defer fmt.Println("defer2")

	fmt.Println("main_end")

	fmt.Println(unintuitive())
	SimeFunc()

}
func unintuitive() (value string) {
	defer func() { value = "На самом деле" }() // круглые скобки в конце означают, что функция вызывается
	return "Казалось бы"
}

func SimeFunc() {
	a := "some text"
	defer func(s string) {
		fmt.Println(s)
	}(a)
	a = "another text"
}

func metricTime(start time.Time) {
	// функция Now() возвращает текущее время, а функция Sub возвращает разницу между двумя временными метками
	fmt.Println(time.Now().Sub(start))
}
