package main

import (
	"fmt"
	"time"
)

func main()  {
	type Employee struct {
		ID int
		Name string
		Address string
		Dob time.Time
		Position string
		Salary int
		ManagerID int
	}

	var dilbert Employee

	fmt.Println(dilbert)
}
