package main

import "fmt"

type employee struct {
	name   string
	sex    string // пол
	age    int
	salary int // зарплата
}

func main() {

	firstEmployee := newEmployee("Вася", "М", 25, 1500)
	fmt.Println(firstEmployee.sex)
	fmt.Printf("%+v\n", firstEmployee)
	fmt.Printf("Працівник %d \n", firstEmployee.salary)
}

// функція конструктор
func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}
