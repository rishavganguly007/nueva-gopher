package main

import (
	"fmt"

	"math/rand"

	"time"
)

func  main()  {
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1
	fmt.Println("dice : ", dice)

	switch dice {
	case 1: 
		fmt.Println("open")
	case 2:
		fmt.Println("move spot by ", dice)
	case 3:
		fmt.Println("move spot by ", dice)
		fallthrough
	case 4:
		fmt.Println("move spot by ", dice)
	case 5:
		fmt.Println("move spot by ", dice)
	case 6:
		fmt.Println("move spot by ", dice)
	}
}