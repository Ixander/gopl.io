package main

import (
	"bytes"
	"fmt"
	"net"
)

var (
	// specify the range of trusted IP addresses
	start = net.ParseIP("10.192.0.1")
	end   = net.ParseIP("10.207.255.254")
)

func checkIP(ip string) bool {

	//sanity check
	input := net.ParseIP(ip)
	if input.To4() == nil {
		fmt.Printf("%v is not a valid IPv4 address\n", input)
		return false
	}

	if bytes.Compare(input, start) >= 0 && bytes.Compare(input, end) <= 0 {
		fmt.Printf("%v is between %v and %v\n", input, start, end)
		return true
	}
	fmt.Printf("%v is NOT between %v and %v\n", input, start, end)
	return false
}

func main() {

	fmt.Println(checkIP("10.193.111.129"))
	fmt.Println(checkIP("10.197.102.193"))
	fmt.Println(checkIP("10.199.37.1"))
	fmt.Println(checkIP("10.200.107.129"))
	fmt.Println(checkIP("10.200.35.65"))
	fmt.Println(checkIP("10.200.38.193"))
	fmt.Println(checkIP("10.200.46.193"))
	fmt.Println(checkIP("10.201.101.1"))
	fmt.Println(checkIP("10.201.102.1"))
	fmt.Println(checkIP("10.202.230.129"))
	fmt.Println(checkIP("10.203.32.49"))

	fmt.Println(checkIP("216.14.49.185"))
	fmt.Println(checkIP("1.2.3.4"))
	fmt.Println(checkIP("198.162.1.102"))
	fmt.Println(checkIP("1::16"))
}
