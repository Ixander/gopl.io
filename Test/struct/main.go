package main

import (
	"fmt"
	"time"
)

func main() {
	type Employee struct {
		ID        int
		Name      string
		Address   string
		Dob       time.Time
		Position  string
		Salary    int
		ManagerID int
	}

	var dilbert Employee
	dilbert.Salary = 500
	var employeeOfTheMonth *Employee = &dilbert
	fmt.Println(employeeOfTheMonth)
	employeeOfTheMonth.Position += "active member of team"
	fmt.Println(employeeOfTheMonth)
	//dilbert.Position = "middle"
	position := &dilbert.Position
	*position = "Senior " + *position

	fmt.Println(position, *position)

}
