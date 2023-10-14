// resource: https://www.youtube.com/watch?v=FjxK1KJ7iq4

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question-answer' ")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Falied to open: %s\n ", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	line, err := r.ReadAll()
	if err != nil {
		exit("Failed to provide csv")
	}
	problems := parselines(line)
	fmt.Println(problems)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n", &answer)	
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nTime ran out. You Scored %d out of %d.\n ",correct, len(problems))
			return
		case answer := <- answerCh:
			if answer == p.a {
				correct++
			}
		
		
	}
	}
	fmt.Printf("You Scored %d out of %d.\n ",correct, len(problems))
}

func parselines(lines [][]string) []problem{
		ret := make([]problem, len(lines))

		for i, line := range lines {
			ret[i] = problem{
				q: line[0],
				a: strings.TrimSpace(line[1]), // handles the case, if csv file has spaces in answer
			}

		}
		return ret
}

type problem struct {
	q string 
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
