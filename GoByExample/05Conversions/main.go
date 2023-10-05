package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome")
	fmt.Println("rate")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks, ", input)

	numRate, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("numRate: ", numRate + 1)
	}
}