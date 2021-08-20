package main

import "fmt"

type employee struct {
	name   string
	sex    string // пол
	age    int
	salary int // зарплата
}

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
	// передаємо не саму структуру а вказівник на неї &firstEmployee
	setName(&firstEmployee, "Гена")
	fmt.Println(firstEmployee.getInfo())
}

// e це вказівник на тип структури employee
func setName(e *employee, name string) {
	fmt.Println(e)
	fmt.Println(*e)
	e.name = name
}
