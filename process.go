package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func processLine(line string, history []historyLine) Flashcard {
	card := parseFlashcard(line)

	for _, line := range history {
		if card.String() == line.card {
			if line.correct {
				card.bin += 1
			} else {
				card.bin = 1
			}

			if card.bin > 5 {
				card.bin = 5
			}

			card.last_reviewed = line.last_reviewed
		}
	}

	return card
}

func processFile(fi os.FileInfo, history []historyLine, config Config) []Flashcard {
	cards := make([]Flashcard, 0)

	if fi.Name()[len(fi.Name())-len(config.suffix):] != config.suffix {
		return cards
	}

	contents, err := ioutil.ReadFile(config.deck_dir + "/" + fi.Name())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		if len(line) > 7 && line[:7] == "#flash|" {
			card := processLine(line, history)
			cards = append(cards, card)
		}
	}

	return cards
}

func processDirectory(history []historyLine, config Config) []Flashcard {
	cards := make([]Flashcard, 0)

	dir, err := os.Open(config.deck_dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)

		restoreScreen()
		os.Exit(1)
	}

	for _, fi := range fileInfos {
		if fi.Mode().IsRegular() {
			newCards := processFile(fi, history, config)
			cards = append(cards, newCards...)
		}
	}

	return cards
}
