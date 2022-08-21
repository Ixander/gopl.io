package main

import "fmt"

type Person struct {
	Name    string
	Surname string
	Age     int
}
type Item struct {
	NoOption   string
	Parameter1 string
	Parameter2 int
	Parameter3 string
}
type option func(*Item)

func main() {
	// с параметрами по умолчанию
	item1 := NewItem()
	fmt.Println(item1)
	// с применением одной опции
	item2 := NewItem(Option2(70))
	fmt.Println(item2)
	// или двух
	item3 := NewItem(Option1("unusual"), Option2(99))
	fmt.Println(item3)
	// опции можно заявлять в разном порядке
	item4 := NewItem(Option2(88), Option1("rare"))
	fmt.Println(item4)

	item5 := NewItem(Option2(88), Option1("rare"), Option3("TEST"))
	fmt.Println(item5)
}
func Option1(option1 string) option {
	return func(i *Item) {
		i.Parameter1 = option1
	}
}
func Option2(option2 int) option {
	return func(i *Item) {
		i.Parameter2 = option2
	}
}
func Option3(option3 string) option {
	return func(i *Item) {
		i.Parameter3 = option3
	}
}

func NewItem(opts ...option) *Item {
	// инициализируем типовыми значениями
	i := &Item{
		NoOption:   "usual",
		Parameter1: "default",
		Parameter2: 42,
	}
	// применяем опции в том порядке, в котором они были заявлены
	for _, opt := range opts {
		opt(i)
	}
	return i
}
