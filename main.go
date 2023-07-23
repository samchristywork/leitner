package main

import (
	"fmt"
	"os"
	"sort"
)

func usage() {
	fmt.Println("Usage: flash [quiz|list|bins] [deck]")
	os.Exit(0)
}

func processArgs(args []string, decks map[string][]Flashcard) {
	if len(args) == 1 {
		if args[0] == "quiz" {
			alternateScreen()
			results := quiz(decks)

			restoreScreen()
			printQuizResults(results)
			os.Exit(0)
		}
	}

	if len(args) == 1 {
		if args[0] == "bins" {
			fmt.Println("Bins:")
			fmt.Println("")

			printBins(decks)

			os.Exit(0)
		}
	}

	if len(args) == 1 {
		if args[0] == "list" {
			blue()
			fmt.Println("Decks:")
			reset()
			fmt.Println("")

			sortedDecks := []string{}

			for deck := range decks {
				sortedDecks = append(sortedDecks, deck)
			}

			sort.Strings(sortedDecks)

			for _, deck := range sortedDecks {
				fmt.Println("•", deck)
			}

			os.Exit(0)
		}

		usage()
	}

	if len(args) == 2 {
		if args[0] == "list" {
			if _, ok := decks[args[1]]; !ok {
				fmt.Println("Deck not found")
			}

			for _, card := range decks[args[1]] {
				fmt.Println(card)
			}
			os.Exit(0)
		}

		usage()
	}

	usage()
}

func main() {
	history := readHistory()

	cards := processDirectory(history, "./")

	decks := make(map[string][]Flashcard)
	for _, card := range cards {
		decks[card.deck] = append(decks[card.deck], card)
	}

	processArgs(os.Args[1:], decks)
}
