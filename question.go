package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func writeQuestionToFile(card Flashcard, currentQuestion int, totalQuestions int, config Config) {
	question := fmt.Sprintf("\n%s", card.front)

	content := []byte("Question " + fmt.Sprint(currentQuestion) + "/" + fmt.Sprint(totalQuestions) + " (" + card.deck + ")\n" + question + "\n---\n")

	err := ioutil.WriteFile(config.question_filename, []byte(content), 0644)
	if err != nil {
		fmt.Println(err)

		restoreScreen()
		os.Exit(1)
	}
}

func appendResultToFile(card Flashcard, result bool, config Config) {
	f, err := os.OpenFile(config.history_filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)

		restoreScreen()
		os.Exit(1)
	}

	defer f.Close()

	currentTime := time.Now().Unix()

	if result {
		if _, err := f.WriteString("quiz	" + card.String() + "	correct	" + fmt.Sprint(currentTime) + "\n"); err != nil {
			fmt.Println(err)

			restoreScreen()
			os.Exit(1)
		}
	} else {
		if _, err := f.WriteString("quiz	" + card.String() + "	incorrect	" + fmt.Sprint(currentTime) + "\n"); err != nil {
			fmt.Println(err)

			restoreScreen()
			os.Exit(1)
		}
	}
}

func askQuestion(card Flashcard, currentQuestion int, totalQuestions int, config Config) bool {
	result := true

	for {
		writeQuestionToFile(card, currentQuestion, totalQuestions, config)

		restoreScreen()
		execProgram("nvim", []string{"-c", ":normal! Go", "--clean", config.question_filename})
		alternateScreen()

		content, err := ioutil.ReadFile(config.question_filename)
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
		fmt.Println("")
		reset()

		blue()
		fmt.Println("Difference:")
		reset()

		r := tokenizeString(parts[1])
		c := tokenizeString(card.back)
		lcs := lcs(r, c)

		d := diff(r, c, lcs)
		printDiff(d)

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
			fmt.Println("Study for " + fmt.Sprint(config.incorrect_review_time) + " seconds and try again.")
			time.Sleep(time.Duration(config.incorrect_review_time) * time.Second)
		}
	}

	appendResultToFile(card, result, config)
	return result
}
