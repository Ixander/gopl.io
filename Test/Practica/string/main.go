package main

import (
	"fmt"
	"time"
)

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

	now := time.Now()

	fmt.Println(now)
	//ver = v0.0.1 id = 0 pi = 3.1415.
	// определите переменные ver, id, pi
	ver := "v0.0.1"
	id := 0
	pi := 3.1415

	fmt.Println("ver =", ver, "id =", id, "pi =", pi)

}
