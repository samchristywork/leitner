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

func processArgs(args []string, cards []Flashcard, config Config) {
	if len(args) == 1 {
		if args[0] == "quiz" {
			alternateScreen()
			results := quiz(cards, config)

			restoreScreen()
			printQuizResults(results)

			os.Exit(0)
		}

		if args[0] == "dump" {
			dumpConfig(config)

			os.Exit(0)
		}

		if args[0] == "bins" {
			fmt.Println("Bins:")
			fmt.Println("")

			printBins(cards, "")

			os.Exit(0)
		}

		if args[0] == "list" {
			blue()
			fmt.Println("Decks:")
			reset()
			fmt.Println("")

			printDecks(cards)

			os.Exit(0)
		}

		usage()
	}

	if len(args) == 2 {
		if args[0] == "list" {

			for _, card := range cards {
				if card.deck == args[1] {
					fmt.Println(card)
				}
			}
			os.Exit(0)
		}

		usage()
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
