// Examples taken from: https://gobyexample.com

/*
 package main: This line declares that this Go file belongs to the main package.
 In Go, the main package is a special package,
 and a Go program always starts executing from the main function within the main package.

 import "fmt": This line imports the "fmt" package, which is a standard library package in Go.
 The "fmt" package provides functions for formatted input and output. In this program, we are using fmt.
 Println to print a message to the console.

 func main() { ... }: This is the main function. It's the entry point for the program and
 where the program execution begins.
*/

package main

import (
	"fmt"
	"Values"

)

func main() {
	fmt.Println("heello-worrld")
	Values.Values()
}
