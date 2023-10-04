package main
// https://go.dev/ref/spec
import "fmt"

func main() {
	var username string = "kewl"
	fmt.Println(username)
	fmt.Printf("Variable is of Type: %T \n", username)

	var isVal bool = true
	fmt.Printf(" variable Type: %T",isVal)

	// No var style
	num := 3000
	fmt.Println(num)
}