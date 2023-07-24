package main

import (
	"fmt"
)

type QuizScore struct {
	correct   int
	incorrect int
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

	printBins(decks, selected_deck)

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
