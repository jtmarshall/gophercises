package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// Question struct to hold a question and answer
type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func main() {
	// Read filename from command line flag, otherwise default to "problems.csv"
	csvFlag := flag.String("csv", "problems.csv", "CSV filename for questions (default is \"problems.csv\"")
	// Quiz timer flag
	timerFlag := flag.Int("timer", 30, "Set time limit of quiz (default is 30sec)")
	flag.Parse()
	// open questions file, and read csv
	csvFile, _ := os.Open(*csvFlag)
	r := csv.NewReader(bufio.NewReader(csvFile))

	// initialize list of questions
	var questions []Question

	// Read csv line by line appending to the list of questions
	for {
		// read line
		line, err := r.Read()
		// break if we get to end of file
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// append question and answer into list of questions
		questions = append(questions, Question{
			Question: line[0],
			Answer:   line[1],
		})
	}

	// Store user's score
	score := 0

	// Ask user to hit Enter to start the Quiz
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hit ENTER to Start")
	_, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Start timer
	timer := time.NewTimer(time.Duration(*timerFlag) * time.Second)
	go func() {
		<-timer.C
		fmt.Printf("\nTimes up!\nYou scored %d out of %d\n", score, len(questions))
		os.Exit(0)
	}()

	// Quiz loop through list of questions asking them to user
	for i, quest := range questions {
		// ask the question
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Problem #%d: %s = ", i+1, quest.Question)

		// get user input
		txt, _ := reader.ReadString('\n')

		// increment score if correct; (trim extra whitespace)
		if strings.TrimSpace(txt) == strings.TrimSpace(quest.Answer) {
			score++
		}
	}

	// Print out user score and number of questions at the end
	fmt.Printf("You scored %d out of %d\n", score, len(questions))
}
