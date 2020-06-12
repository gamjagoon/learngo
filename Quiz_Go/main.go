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
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,anwer'")
	timeLimit := flag.Int("Limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

  var correct int32 = 0
  problemloop:
    for i, p := range problems {
      fmt.Printf("Proble #%d: %s = ", i+1, p.question)
      anserCh := make(chan string)
      go func ()  {
        var answer string
        fmt.Scanf("%s\n", &answer)
        anserCh <- answer
      }()

      select {
      case <-timer.C:
        fmt.Println()
        break problemloop
      case answer := <- anserCh:
        if answer == p.answer {
          correct++
        }
      }
    }

    fmt.Printf("\nYou scored %d out of %d \n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.Trim(line[1], " "),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
