package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type QuizScore struct {
	correct   int
	incorrect int
}

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
		green()
		fmt.Println("Correct answer:")
		reset()
		fmt.Println(card.back)
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

func administerQuiz(questions []Flashcard, config Config) QuizScore {
	score := QuizScore{0, 0}

	for n, card := range questions {
		result := askQuestion(card, n+1, len(questions), config)

		if result {
			score.correct++
		} else {
			score.incorrect++
		}
	}

	return score
}

func quiz(decks map[string][]Flashcard, config Config) QuizScore {
	clearScreen()

	var selected_num int
	var selected_deck string
	var selected_bin int

	blue()
	fmt.Println("Filters:")
	reset()
	fmt.Println("")

	// Filter by number of cards
	blue()
	fmt.Println("How many cards? (0 for all)")
	reset()
	fmt.Println("")
	fmt.Print("❯ ")

	fmt.Scanf("%d", &selected_num)

	// Filter by deck
	blue()
	fmt.Println("")
	fmt.Println("Which deck? (Empty for all decks)")
	reset()
	fmt.Println("")

	for deck := range decks {
		fmt.Println("•", deck)
	}

	fmt.Println("")
	fmt.Print("❯ ")

	fmt.Scanf("%s", &selected_deck)

	// Filter by bin
	blue()
	fmt.Println("")
	fmt.Println("Which bin? (Empty for all bins)")
	reset()
	fmt.Println("")

	printBins(decks)

	fmt.Println("")
	fmt.Print("❯ ")

	fmt.Scanf("%d", &selected_bin)

	questions := []Flashcard{}
	for _, deck := range decks {
		for _, card := range deck {
			if card.bin == uint32(selected_bin) || selected_bin == 0 {
				if card.deck == selected_deck || selected_deck == "" {
					questions = append(questions, card)
				}
			}
		}
	}

	if selected_num == 0 {
		selected_num = len(questions)
	}

	if selected_num > len(questions) {
		selected_num = len(questions)
	}

	randomizedQuestions := shuffle(questions, selected_num)

	score := administerQuiz(randomizedQuestions, config)
	return score
}

func printQuizResults(results QuizScore) {
	correct := results.correct
	incorrect := results.incorrect
	total := results.correct + results.incorrect
	percent := float64(correct) / float64(total) * 100

	blue()
	fmt.Println("Quiz Results:")
	fmt.Println("─────────────")
	reset()
	green()
	fmt.Println("Correct:", correct)
	reset()
	red()
	fmt.Println("Incorrect:", incorrect)
	reset()
	fmt.Println("Total:", total)
	fmt.Printf("Percent: %.2f%%\n", percent)
}
