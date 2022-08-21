package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// metricTimeCall — функция-обёртка для замера времени
func metricTimeCall(f func(string)) func(string) {
	fmt.Println("step2")
	return func(s string) {
		fmt.Println("step4")
		start := time.Now() // получаем текущее время
		f(s)
		fmt.Println("Execution time: ", time.Now().Sub(start)) // получаем интервал времени как разницу между двумя временными метками
		fmt.Println("step7")
	}
}

// countCall — функция-обёртка для подсчёта вызовов
func countCall(f func(string)) func(string) {
	cnt := 0
	// получаем имя функции. Подробнее об этом вызове будет рассказано в следующем курсе
	funcname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return func(s string) {
		cnt++
		fmt.Printf("Функция %s вызвана %d раз\n", funcname, cnt)
		f(s)
	}
}
func myprint(s string) {
	fmt.Println("step5")
	fmt.Println(s)
	fmt.Println("step6")
}
func main() {

	/*fmt.Println("step1")
	countTime := metricTimeCall(myprint)
	fmt.Println("step3")
	countTime("TEST")*/

	countedPrint := countCall(myprint)
	countedPrint("Hello world")
	countedPrint("Hi")

	// обратите внимание, что мы оборачиваем уже обёрнутую функцию, а значение счётчика вызовов при этом сохраняется
	//countAndMetricPrint := metricTimeCall(countedPrint)
	//countAndMetricPrint("Hello")
	//countAndMetricPrint("World")

}
