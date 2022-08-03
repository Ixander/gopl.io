package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Получаем читателя пользовательского ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	cnt := 0
	for {
		fmt.Print("-> ")
		// Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		f(&cnt)

		fmt.Printf("User input %d lines\n", cnt)
	}
}

func f(p *int) {
	*p = *p + 1
}

//Перед вами неполный код программы, которая считает, сколько строк пользователь ввёл в консоль,
//и после ввода каждой новой строки выводит общее количество на экран. Напишите реализацию функции f.
