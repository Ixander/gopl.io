package main

import "fmt"

//тип структури
type employee struct {
	name   string
	sex    string // пол
	age    int
	salary int // зарплата
}

//конструктор
func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}

func (e employee) getInfo() string {
	return fmt.Sprintf("Сотрудник: %s\nПол: %s\nВозраст: %d\nЗарплата: %d", e.name, e.sex, e.age, e.salary)
}

func main() {
	firstEmployee := newEmployee("Вася", "М", 25, 1500)
	fmt.Println(firstEmployee.getInfo())
}
