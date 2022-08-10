package main

import "fmt"

func main() {

	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("type: %T \n", weekTempArr)

	workDaysSlice := weekTempArr[:5]
	weekendSlice := weekTempArr[5:]
	fromTuesdayToThursDaySlice := weekTempArr[1:4]
	weekTempSlice := weekTempArr[:]

	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))                                        // [1 2 3 4 5] 5 7
	fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))                                           // [6 7] 2 2
	fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6
	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))                                        // [1 2 3 4 5 6 7] 7 7
	fmt.Println("----------------------")
	s := make([]int, 4, 7) // [0 0 0 0], len = 4 cap = 7

	slice1 := append(s[:2], 2, 3, 4)
	fmt.Println(s, slice1) // [0 0 2 3] [0 0 2 3 4]
	fmt.Println("----------------------")
	fmt.Println(slice1)
	testSlice(&slice1)
	fmt.Println(slice1)

	fmt.Println("----------------------")
	s = []int{1, 2, 3, 4, 5}
	i := 2

	if len(s) != 0 && i < len(s)-1 { // защищаемся от паники
		s = append(s[:i], s[i+1:]...)
	}
	fmt.Println(s) // [1 2 4 5]
	fmt.Printf("%q\n", s)

}

func testSlice(s *[]int) {

	*s = append(*s, []int{7, 8}...)

}
