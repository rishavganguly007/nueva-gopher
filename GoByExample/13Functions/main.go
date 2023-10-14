package main

import "fmt"

func  main()  {
	fmt.Println("in func()")
	greet()
	fmt.Println("Add: ", add(1,2))
	fmt.Println("Add-multi: ", addMulti(1,2, 3,4))

}

func add(val1 int, val2 int) int{
	return val1 + val2
}

func addMulti(value ...int) int{
	total := 0

	for _, val := range value{
		total += val
	}
	 return total
}

func greet() {
	fmt.Println("in greet()")
}