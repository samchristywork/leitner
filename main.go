package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage: flash [quiz|list|bins] [deck]")
	os.Exit(0)
}

func processArgs(args []string, cards []Flashcard, config Config) {
	if len(args) == 1 && args[0] == "quiz" {
		alternateScreen()
		results := quiz(cards, config)

		restoreScreen()
		printQuizResults(results)

		os.Exit(0)
	} else if len(args) == 2 && args[0] == "list" && args[1] == "bins" {
		fmt.Println("Bins:")
		fmt.Println("")

		printBins(cards, "")

		os.Exit(0)
	} else if len(args) == 2 && args[0] == "list" && args[1] == "config" {
		dumpConfig(config)

		os.Exit(0)
	} else if len(args) == 3 && args[0] == "list" && args[1] == "deck" {
		found := false

		for _, card := range cards {
			if card.deck == args[2] {
				fmt.Println(card)
				found = true
			}
		}

		if !found {
			fmt.Println("Deck not found")
		}

		os.Exit(0)
	} else if len(args) == 2 && args[0] == "list" && args[1] == "decks" {
		blue()
		fmt.Println("Decks:")
		reset()
		fmt.Println("")

		printDecks(cards)

		os.Exit(0)
	} else if len(args) == 2 && args[0] == "list" && args[1] == "cards" {
		for _, card := range cards {
			fmt.Println(card)
		}

		os.Exit(0)
	}

	usage()
}

func main() {
	home := os.Getenv("HOME")

	configFilename := home + "/.config/leitner/config"

	config := readConfigFile(configFilename)
	history := readHistory(config.historyFile)
	cards := processDirectory(history, config.deckdir)

	processArgs(os.Args[1:], cards, config)
}
