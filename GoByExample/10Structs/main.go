package main

import "fmt"

func  main()  {
 aUser := User{"name", true, 22}
 fmt.Printf("A user detals : %+v\n",aUser)
}

type User struct{
	name string
	status bool
	age int

}