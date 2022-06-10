package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	var (
	//out bytes.Buffer
	//input string
	)

	fmt.Println("Hello")

	/*c := exec.Command("ll")
	c.Stdout = os.Stdout
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}*/

	//cmd := exec.Command("tr", "a-z", "A-Z")
	cmd := exec.Command("ls", "-lah")

	//fmt.Scanln(&input)

	//cmd.Stdin = strings.NewReader(input)
	//cmd.Stdout = &out
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("in all caps: %q\n", out.String())

}
