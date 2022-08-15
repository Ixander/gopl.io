package main

import "fmt"

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
}
