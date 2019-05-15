package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// Question struct to hold a question and answer
type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func main() {
	filename := "problems.csv"
	// open questions file, and read csv
	csvFile, _ := os.Open(filename)
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

	// questionsJSON, _ := json.Marshal(questions)
	fmt.Println(questions)

	// Loop through list of questions asking them to user
	for _, quest := range questions {
		// ask the question
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(quest.Question)

		// wait for user input
		txt, _ := reader.ReadString('\n')
	}
}
