// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 179.

// The sleep program sleeps for a specified period of time.
package main

import (
	"flag"
	"fmt"
	"time"
)

//!+sleep
var period = flag.Duration("period", 1*time.Second, "sleep period")
var test = flag.String("test", "not set", "test")

func main() {

	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)

	fmt.Println(*test)
}

//!-sleep
