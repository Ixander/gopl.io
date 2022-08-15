package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	type MyMap map[string]string

	var m1 MyMap
	m1 = make(MyMap, 5)

	// объект готов
	m1["foo"] = "bar"

	v, ok := m1["foo1"]

	fmt.Println(v, ok)

	fmt.Println(m1, len(m1))

	sentence := "Πολύ λίγα πράγματα συμβαίνουν στο σωστό χρόνο, και τα υπόλοιπα δεν συμβαίνουν καθόλου"
	// инициализируем map
	// ключами будут знаки, а значениями — количество вхождений
	frequency := make(map[rune]int)
	// пройдёмся по знакам в предложении
	for _, v := range sentence {
		frequency[v]++ // встреченному знаку увеличиваем частоту
	}
	// печатаем
	for k, v := range frequency {
		fmt.Printf("Знак %c встречается %d раз \n", k, v)
	}

	groceries := map[string]int{
		"хліб":     50,
		"молоко":   100,
		"масло":    200,
		"ковбаса":  500,
		"сіль":     20,
		"огірки":   200,
		"сир":      600,
		"шинка":    700,
		"буженина": 900,
		"помідори": 250,
		"риба":     300,
		"хамон":    1500,
	}
	order := []string{"хліб", "буженина", "сир", "огірки"}

	sum := 0
	for k, v := range groceries {
		if v > 500 {
			fmt.Println(k)
		}
		for o := range order {
			if k == order[o] {
				sum += v
			}
		}

	}
	fmt.Println(sum)

	//Дан массив целых чисел A и целое число k.
	//Нужно найти и вывести индексы пары чисел, сумма которых равна k.
	//Если таких чисел нет, то вернуть пустой слайс. Индексы можно вернуть в любом порядке

	var k int
	fmt.Print("Input k: ")
	fmt.Scanln(&k)

	slice := make([]int, 0)

	fmt.Println("Input mass: ")
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		inp, _ := strconv.Atoi(input.Text())
		slice = append(slice, inp)
	}
	fmt.Println(slice)
	fmt.Println(findIndex(slice, k))
	fmt.Println(find(slice, k))

}

func findIndex(mass []int, k int) []int {

	res := make([]int, 0)
	data := make(map[int]int)
	for k, v := range mass {
		data[k] = v
	}
	fmt.Println(data)

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if k-data[i] == data[j] && data[i] != data[j] {
				res = append(res, i)
			}
		}
		fmt.Println(data[i])
	}
	return res
}

func find(arr []int, k int) []int {
	// Создадим пустую map
	m := make(map[int]int)
	// будем складывать в неё индексы массива, а в качестве ключей использовать само значение
	for i, a := range arr {
		if j, ok := m[k-a]; ok { // если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k и мы нашли, то что нужно
			return []int{i, j}
		}
		// если искомого значения нет, то добавляем текущий индекс и значение в map
		m[a] = i
	}
	// не нашли пары подходящих чисел
	return nil
}

// как можно заметить, алгоритм пройдётся по массиву всего один раз
// если бы мы искали подходящее значение каждый раз через перебор массива,
