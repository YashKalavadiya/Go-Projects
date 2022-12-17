package main

import (
	"flag"
	"os"
	"fmt"
	"encoding/csv"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in format of question, ans")

	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")


	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv File %s", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("Error in %s", *csvFilename))
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, p := range(problems) {

		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)

		ansChan := make(chan string)
		go func() {
			var ans string

			fmt.Scanf("%s\n", &ans)
			
			ansChan <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("Your scored %d out of %d. \n", correct, len(problems))
			return
		case ans := <-ansChan:
			if(ans == p.a) {
				correct++
			}
		}
		
	}
}

	
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range(lines) {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}


func exit(msg string) {
	fmt.Println(msg);
	os.Exit(1)
}
