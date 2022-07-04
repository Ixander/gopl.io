package main

import "fmt"

func foo() {
	// паникуем
	panic("unexpected! 777")
}

//...
// выполняется после срабатывания паники
func main() {
	defer func() {
		if r := recover(); r != nil {
			// обработка паники, в переменной r будет лежать строка "unexpected!"
			fmt.Println(r)
		}
	}()
	// внутри foo срабатывает паника
	foo()
}
