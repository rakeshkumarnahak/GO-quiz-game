package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type questionAnswers struct {
	question string
	answer   string
}

func main() {
	//cmd flag feature
	strPtr := flag.String("filename", "problems.csv", "file found")
	flag.Parse()

	var commandIs string

	if *strPtr != "problems.csv" {
		commandIs = *strPtr + ".csv"
	} else {
		commandIs = *strPtr
	}

	fmt.Println("Welcome to quiz game")

	//Opening the csv file got from the cmd (or the default value) using os.Open()
	file, err := os.Open(commandIs)
	checkNilErr(err)

	//Reading the opened csv file using csv.NewReader()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	checkNilErr(err)

	//Exporting the data in csv file into a struct
	questionAnswerList := createQuestionsAnswers(data)

	//Checking the exported questionAnswer data
	// fmt.Printf("%+v\n", questionAnswerList)

	fmt.Println("Let's start the game by typing answers of all the following questions.") //First Messege to the user after starting the game

	var totalScore int

	// reader := bufio.NewReader(os.Stdin) //reader to read user input

	//Printing the question one by one and measuring the score
	for i, QApair := range questionAnswerList {
		fmt.Printf("%v. %s = ", i+1, QApair.question)

		var userinput int64
		//Taking the userinput
		fmt.Scanln(&userinput)
		correctAnswer, _ := strconv.ParseInt(QApair.answer, 10, 64)

		//Checking the answer
		if userinput == correctAnswer {
			totalScore += 1
		} else if userinput != correctAnswer && totalScore >= 0 {
			totalScore -= 1
		}
	}

	fmt.Printf("Your Final Score is %v", totalScore)
}

// function to add quetion and answers into the struct from the csv file
func createQuestionsAnswers(data [][]string) []questionAnswers {
	var questionAnswerList []questionAnswers

	for _, line := range data {
		var rec questionAnswers
		for j, field := range line {
			if j == 0 {
				rec.question = field
			} else if j == 1 {
				rec.answer = field
			}
		}
		questionAnswerList = append(questionAnswerList, rec)
	}
	return questionAnswerList
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
