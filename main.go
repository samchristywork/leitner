package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " <command> [<args>]")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  quiz")
	fmt.Println("  list bins")
	fmt.Println("  list config")
	fmt.Println("  list deck <deck>")
	fmt.Println("  list decks")
	fmt.Println("  list cards")
	os.Exit(0)
}

func processArgs(cards []Flashcard, config Config) {
	args := os.Args

	if len(args) == 2 && args[1] == "quiz" {
		alternateScreen()
		results := quiz(cards, config)

		restoreScreen()
		printQuizResults(results)

		os.Exit(0)
	} else if len(args) == 3 && args[1] == "list" && args[2] == "bins" {
		fmt.Println("Bins:")
		fmt.Println("")

		printBins(cards, "")

		os.Exit(0)
	} else if len(args) == 3 && args[1] == "list" && args[2] == "config" {
		dumpConfig(config)

		os.Exit(0)
	} else if len(args) == 4 && args[1] == "list" && args[2] == "deck" {
		found := false

		for _, card := range cards {
			if card.deck == args[3] {
				fmt.Println(card)
				found = true
			}
		}

		if !found {
			fmt.Println("Deck not found")
		}

		os.Exit(0)
	} else if len(args) == 3 && args[1] == "list" && args[2] == "decks" {
		blue()
		fmt.Println("Decks:")
		reset()
		fmt.Println("")

		printDecks(cards)

		os.Exit(0)
	} else if len(args) == 3 && args[1] == "list" && args[2] == "cards" {
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
	history := readHistory(config.history_filename)
	cards := processDirectory(history, config)

	processArgs(cards, config)
}
