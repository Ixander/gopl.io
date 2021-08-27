package main

import "fmt"

type IntList struct {
	Value int
	Tail  *IntList
}

func main() {
	i := IntList{
		Value: 5,
		Tail:  nil,
	}

	j := IntList{
		Value: 3,
		Tail:  &i,
	}

	fmt.Println(i)
	fmt.Println(i.Sum())
	fmt.Println(j.Sum())
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
