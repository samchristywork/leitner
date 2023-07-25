package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"time"
)

func hashString(s string) uint32 {
	var hash uint32 = 5381
	for _, c := range s {
		hash = ((hash << 5) + hash) + uint32(c)
	}
	return hash
}

func shuffle(cards []Flashcard, num int) []Flashcard {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards[:num]
}

func execProgram(program string, args []string) {
	cmd := exec.Command(program, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func readLine() string {
	var answer string
	fmt.Scanln(&answer)
	return answer
}

func alternateScreen() {
	fmt.Print("\033[?1049h")
}

func restoreScreen() {
	fmt.Print("\033[?1049l")
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func blue() {
	fmt.Print("\033[34m")
}

func green() {
	fmt.Print("\033[32m")
}

func red() {
	fmt.Print("\033[31m")
}

func reset() {
	fmt.Print("\033[0m")
}

func die(status int) {
	restoreScreen()
	os.Exit(status)
}

func binHistogram(cards []Flashcard, deckFilter string) map[uint32]int {
	bins := make(map[uint32]int)

	for _, card := range cards {
		if card.deck == deckFilter || deckFilter == "" {
			bin := card.bin

			if _, ok := bins[bin]; !ok {
				bins[bin] = 1
			} else {
				bins[bin]++
			}
		}
	}

	return bins
}

func printBins(cards []Flashcard, deckFilter string) {
	histogram := binHistogram(cards, deckFilter)

	keys := []uint32{}
	for key := range histogram {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, key := range keys {
		fmt.Printf("• Bin %d: %d\n", key, histogram[key])
	}
}
