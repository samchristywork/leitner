package main

import (
	"fmt"
	"time"
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

func askTime(cards []Flashcard) int64 {
	var selected_time int64

	blue()
	fmt.Println("")
	fmt.Println("How stale (in days) should the cards be? (0 for no filter)")
	reset()
	fmt.Println("")

	printDaysSince(cards)

	fmt.Println("")
	fmt.Print("❯ ")

	fmt.Scanf("%d", &selected_time)

	return selected_time
}

func isInBlacklist(deck string, config Config) bool {
	for _, blacklist := range config.blacklist {
		if deck == blacklist {
			return true
		}
	}

	return false
}

func quiz(cards []Flashcard, config Config) QuizScore {
	clearScreen()

	now := time.Now().Unix()

	blue()
	fmt.Println("Filters:")
	reset()
	fmt.Println("")

	filtered_cards := []Flashcard{}

	for _, card := range cards {
		not_in_blacklist := !isInBlacklist(card.deck, config)

		if not_in_blacklist {
			filtered_cards = append(filtered_cards, card)
		}
	}

	cards = filtered_cards

	selected_num := askNumCards()
	selected_deck := askDeck(cards)

	filtered_cards = []Flashcard{}

	for _, card := range cards {
		proper_deck := card.deck == selected_deck ||
			selected_deck == ""

		if proper_deck {
			filtered_cards = append(filtered_cards, card)
		}
	}

	cards = filtered_cards

	selected_bin := askBin(cards, selected_deck)

	filtered_cards = []Flashcard{}

	for _, card := range cards {
		proper_bin := card.bin == uint32(selected_bin) ||
			selected_bin == 0

		if proper_bin {
			filtered_cards = append(filtered_cards, card)
		}
	}

	cards = filtered_cards

	selected_time := askTime(cards)

	filtered_cards = []Flashcard{}

	for _, card := range cards {
		proper_time := selected_time == 0 ||
			card.last_correct == 0 ||
			int64(now)-card.last_correct > selected_time*24*60*60

		if proper_time {
			filtered_cards = append(filtered_cards, card)
		}
	}

	cards = filtered_cards

	if selected_num == 0 {
		selected_num = len(cards)
	}

	if selected_num > len(cards) {
		selected_num = len(cards)
	}

	randomizedQuestions := shuffle(cards, selected_num)

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
