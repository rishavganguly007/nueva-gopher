package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("give val")

	// comma ok OR err, this is how err is handled in GO
	input, _:= reader.ReadString('\n')
	fmt.Println("Your val ", input)
}