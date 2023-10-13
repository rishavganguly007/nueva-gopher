package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question-answer' ")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Falied to open: %s\n ", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	line, err := r.ReadAll()
	if err != nil {
		exit("Faile to provide csv")
	}
	fmt.Println(line)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
