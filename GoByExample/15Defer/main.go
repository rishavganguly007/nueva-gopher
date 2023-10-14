package main

import "fmt"

func main() {
	defer fmt.Println("2")
	defer fmt.Println("1")

	// move the excecution line just before the end of the function and excutes
	// all the defer s in LIFO order
}
