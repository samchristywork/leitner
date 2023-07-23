package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func processFile(fi os.FileInfo, history []historyLine, dir string) []Flashcard {
	cards := make([]Flashcard, 0)

	if fi.Name()[len(fi.Name())-3:] == ".dm" {
		contents, err := ioutil.ReadFile(dir + fi.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lines := strings.Split(string(contents), "\n")
		for _, line := range lines {
			if len(line) > 7 && line[:7] == "#flash|" {
				card := parseFlashcard(line, history)

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
					}
				}

				cards = append(cards, card)
			}
		}
	}

	return cards
}

func processDirectory(history []historyLine, dirname string) []Flashcard {
	cards := make([]Flashcard, 0)

	dir, err := os.Open(dirname)
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
			newCards := processFile(fi, history, dir.Name()+"/")
			cards = append(cards, newCards...)
		}
	}

	return cards
}
