package main

import "fmt"

func  main()  {
	var lst [4]string // len of array MUST be provided

	lst[0] = "a1"
	lst[1] = "a2"
	lst[3] = "a4"
	
	fmt.Println(lst) // Output: [a1 a2  a4], alse len(lst) = 4

}