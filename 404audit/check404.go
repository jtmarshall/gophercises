package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Read filename from command line flag, otherwise default to "404.csv"
	csvFlag := flag.String("csv", "404.csv", "CSV filename for url's to check")
	flag.Parse()
	// open questions file, and read csv
	csvFile, _ := os.Open(*csvFlag)
	r := csv.NewReader(bufio.NewReader(csvFile))

	// Creating csv writer
	wFile, err := os.Create("404result.csv")
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer wFile.Close()
	writer := csv.NewWriter(wFile)

	// Read csv line by line processing each url
	for {
		// read line
		line, err := r.Read()
		// break if we get to end of file
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// check status of url read in
		statusCode := getStatus(line[0])

		wErr := writer.Write([]string{line[0], fmt.Sprint(statusCode)})
		if wErr != nil {
			log.Fatal("ERR writing:", line[0])
		}
		writer.Flush()
	}
}

// pass url in to get the status code for url
func getStatus(url string) int {
	// Start request
	resp, respErr := http.Get(url)
	if respErr != nil {
		fmt.Println(respErr)
		return 0
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode, url)
	return resp.StatusCode
}
