package main

import "fmt"

func  main()  {

	var lst = []string{"a","b","c"}
	fmt.Println("type of list")

	lst = append(lst, "ba")

	fmt.Println(lst)

	fmt.Println("sliced list: ", lst[1:3])

	hghScr := make([]int, 4)
	hghScr[0] = 34
	hghScr[1] = 34
	hghScr[2] = 354
	hghScr[3] = 388
	//hghScr[4] = 389 // this wil give out of bounds

	hghScr = append(hghScr, 54,897,999)

	fmt.Println("High Score: ", hghScr)

	//Removing vals
	var courses = []string{"a","b","c"}
	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses: ", courses)

}