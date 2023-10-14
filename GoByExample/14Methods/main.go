package main

import "fmt"

func  main()  {
 aUser := User{"name", true, 22}
 fmt.Printf("A user detals : %+v\n",aUser)
 aUser.GetStatus()
}

type User struct{
	Name string
	Status bool
	Age int

}

func (u User) GetStatus(){ // 1st letter capital to make it public else it'll be private
	fmt.Println("Status : ", u.Status)
}