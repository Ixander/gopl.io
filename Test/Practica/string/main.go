package main

import "fmt"

func main() {

	var r rune = 161
	text := "Hello"

	text1 := `Hello
			Sasha
		How are you`

	fmt.Println(text)
	fmt.Println(text1)
	fmt.Println(string(r))

	var str string
	str = "Hello, world!"
	println(string(str[0])) //72

}
