package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Config struct {
	deckdir string
	historyFile string
}

func dumpConfig(config Config) {
	fmt.Println("deckdir: " + config.deckdir)
	fmt.Println("history: " + config.historyFile)
}

func readConfigFile(filename string) Config {
	config := Config{}

	home := os.Getenv("HOME")

	// Defaults
	config.historyFile = home + "/.flash_history"
	config.deckdir = "./"

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println("Config file ("+filename+") not found")
		os.Exit(0)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading config file")
		os.Exit(0)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "deckdir:") {
			config.deckdir = strings.TrimSpace(strings.TrimPrefix(line, "deckdir:"))
		}

		if strings.HasPrefix(line, "history:") {
			config.historyFile = strings.TrimSpace(strings.TrimPrefix(line, "history:"))
		}
	}

	return config
}

func usage() {
	fmt.Println("Usage: flash [quiz|list|bins] [deck]")
	os.Exit(0)
}

func processArgs(args []string, decks map[string][]Flashcard, config Config) {
	if len(args) == 1 {
		if args[0] == "quiz" {
			alternateScreen()
			results := quiz(decks, config)

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

			printBins(decks)

			os.Exit(0)
		}

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
	home := os.Getenv("HOME")

	configFilename := home + "/.config/leitner/config"

	config := readConfigFile(configFilename)
	history := readHistory(config.historyFile)
	cards := processDirectory(history, config.deckdir)

	decks := make(map[string][]Flashcard)
	for _, card := range cards {
		decks[card.deck] = append(decks[card.deck], card)
	}

	processArgs(os.Args[1:], decks, config)
}
