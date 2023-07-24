package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func writeQuestionToFile(card Flashcard, currentQuestion int, totalQuestions int) {
	question := fmt.Sprintf("\n%s", card.front)

	content := []byte("Question " + fmt.Sprint(currentQuestion) + "/" + fmt.Sprint(totalQuestions) + " (" + card.deck + ")\n" + question + "\n---\n")

	err := ioutil.WriteFile("/tmp/flashQuestion.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println(err)

		restoreScreen()
		os.Exit(1)
	}
}

func appendResultToFile(card Flashcard, result bool, config Config) {
	f, err := os.OpenFile("/home/sam/.flash_history", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)

		restoreScreen()
		os.Exit(1)
	}

	defer f.Close()

	if result {
		if _, err := f.WriteString(card.String() + "	correct\n"); err != nil {
			fmt.Println(err)

			restoreScreen()
			os.Exit(1)
		}
	} else {
		if _, err := f.WriteString(card.String() + "	incorrect\n"); err != nil {
			fmt.Println(err)

			restoreScreen()
			os.Exit(1)
		}
	}
}

func askQuestion(card Flashcard, currentQuestion int, totalQuestions int, config Config) bool {
	result := true

	for {
		writeQuestionToFile(card, currentQuestion, totalQuestions)

		restoreScreen()
		execProgram("nvim", []string{"-c", ":normal! Go", "--clean", "/tmp/flashQuestion.txt"})
		alternateScreen()

		content, err := ioutil.ReadFile("/tmp/flashQuestion.txt")
		if err != nil {
			fmt.Println(err)

			restoreScreen()
			os.Exit(1)
		}

		clearScreen()
		blue()
		fmt.Println("Question:")
		reset()
		fmt.Println(card.front)
		fmt.Println("")

		blue()
		fmt.Println("You answered:")
		reset()
		parts := strings.Split(string(content), "---\n")
		fmt.Println(parts[1])
		blue()
		fmt.Println("Correct answer:")
		reset()
		green()
		fmt.Println(card.back)
		reset()
		fmt.Println("")

		blue()
		fmt.Println("Is this correct? (y/n)")
		reset()
		fmt.Println("")
		fmt.Print("❯ ")

		answer := readLine()
		if answer == "y" {
			green()
			fmt.Println("Correct!")
			reset()
			break
		} else {
			result = false
			fmt.Println("")
			red()
			fmt.Println("Incorrect!")
			reset()
			fmt.Println("Study for 10 seconds and try again.")
			time.Sleep(10 * time.Second)
		}
	}

	appendResultToFile(card, result, config)
	return result
}
