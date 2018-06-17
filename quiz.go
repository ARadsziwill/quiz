package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// handle csv flag, default is is "problems.csv"
	var problems string
	flag.StringVar(&problems, "csv", "problems.csv", "path to a .csv with problems and answers")
	flag.Parse()

	f, err := os.Open(problems)
	check(err)
	defer f.Close()

	// read the Q&A file
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	quiz(records)

	//fmt.Print(records)

}

func quiz(questions [][]string) {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	var score int
	for _, question := range questions {
		fmt.Printf("%s\n", question[0])
		answer, err := reader.ReadString('\n')
		check(err)
		if strings.EqualFold(strings.TrimSpace(answer), strings.TrimSpace(question[1])) {
			score++
		} else {
			fmt.Printf("Correct answer: %s\n", question[1])
		}
	}
	fmt.Printf("Score: %d/%d\n", score, len(questions))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
