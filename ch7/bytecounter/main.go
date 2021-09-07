// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {

	x := ByteCounter(len(p))
	fmt.Printf("x- %v\n", x)

	var y ByteCounter

	y += ByteCounter(len(p))
	fmt.Printf("y- %v\n", y)

	// c - вказівник на тип ByteCounter
	fmt.Printf("*c- %v\n", *c)
	//*c = *c + ByteCounter(len(p))
	*c += ByteCounter(len(p)) // convert int to ByteCounter

	// *c - значення яке лежить за вказівником с
	//fmt.Println(*c)
	return len(p), nil
}

//!-bytecounter

func main() {
	//!+main
	var c ByteCounter
	_, err := c.Write([]byte("hello"))
	if err != nil {
		return
	}
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	res, err := fmt.Fprintf(&c, "hello, %s", name)

	if err != nil {
		return
	}

	fmt.Println(res)
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main
}
