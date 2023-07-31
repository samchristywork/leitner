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

func askNumCards() int {
	var selected_num int

	blue()
	fmt.Println("How many cards? (0 for all)")
	reset()
	fmt.Println("")
	fmt.Print("❯ ")

	fmt.Scanf("%d", &selected_num)

	return selected_num
}

func askDeck(cards []Flashcard) string {
	var selected_deck string

	blue()
	fmt.Println("")
	fmt.Println("Which deck? (Empty for all decks)")
	reset()
	fmt.Println("")

	printDecks(cards)

	fmt.Println("")
	fmt.Print("❯ ")

	fmt.Scanf("%s", &selected_deck)

	return selected_deck
}

func askBin(cards []Flashcard, selected_deck string) int {
	var selected_bin int

	blue()
	fmt.Println("")
	fmt.Println("Which bin? (Empty for all bins)")
	reset()
	fmt.Println("")

	printBins(cards, selected_deck)

	fmt.Println("")
	fmt.Print("❯ ")

	fmt.Scanf("%d", &selected_bin)

	return selected_bin
}

func quiz(cards []Flashcard, config Config) QuizScore {
	clearScreen()

	blue()
	fmt.Println("Filters:")
	reset()
	fmt.Println("")

	selected_num := askNumCards()
	selected_deck := askDeck(cards)
	selected_bin := askBin(cards, selected_deck)

	questions := []Flashcard{}
	for _, card := range cards {
		if card.bin == uint32(selected_bin) || selected_bin == 0 {
			if card.deck == selected_deck || selected_deck == "" {
				questions = append(questions, card)
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
