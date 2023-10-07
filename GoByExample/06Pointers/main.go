package main

import "fmt"

func main() {
	fmt.Println("hello")

	// var ptr *int

	myNum := 23
	var otr = &myNum
	fmt.Println("adr: ", &otr)
	fmt.Println("val: ", *otr)

	*otr += 2
	fmt.Println("new val: ", myNum)
}