package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)


func readFile(filename string, questions *[]string, answers *[]int) (){

	// open csv file
	csvfile, error := os.Open(filename)
	
	if error != nil {
		log.Fatalln("Can't open CSV file", error)
	}

	// parse the file

	reader := csv.NewReader(csvfile)

	// iterate through the records

	for {
		record, error := reader.Read()

		if error == io.EOF {
			break
		}

		if error != nil {
			log.Fatalln(error)
		}

		*questions = append(*questions, record[0])
		answer, conversionError := strconv.Atoi(record[1])
		if conversionError != nil {
			log.Fatalln("Error in converting string to int", error)
		}
		*answers = append(*answers, answer)
	}
}

func main(){

	questions := make([] string, 0)

	answers := make([] int, 0)

	readFile("problems.csv",&questions, &answers)

	rightAnswers := 0

	for i := 0; i<len(questions);i++{
		fmt.Printf("[Question %d] %s = \n", i+1, questions[i])
		var result int
		fmt.Scanln(&result)

		if result == answers[i]{
			rightAnswers++
			fmt.Println("Correct")
		}else {
			fmt.Println("Wrong")
		}
	}

	fmt.Printf("Final Result: %d/%d\n", rightAnswers, len(questions))
}