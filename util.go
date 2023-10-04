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

	fmt.Println("Bin\tCount")
	fmt.Println("───\t─────")
	for _, key := range keys {
		fmt.Printf("%d\t%d\n", key, histogram[key])
	}
}

func printDecks(cards []Flashcard) {
	decks := make(map[string]int)

	for _, card := range cards {
		if _, ok := decks[card.deck]; !ok {
			decks[card.deck] = 1
		} else {
			decks[card.deck]++
		}
	}

	keys := []string{}

	for key := range decks {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("• %d\t%s\n", decks[key], key)
	}
}

func printDaysSince(cards []Flashcard) {
	daysSinceLastReviewed := make(map[int]int)

	for _, card := range cards {
		now := time.Now().Unix()

		days := (now - card.last_correct) / (60 * 60 * 24)

		if card.last_correct == 0 {
			days = -1
		}

		daysSinceLastReviewed[int(days)]++
	}

	keys := []int{}

	for key := range daysSinceLastReviewed {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	fmt.Println("Days\tCount")
	fmt.Println("────\t─────")
	for _, days := range keys {
		if days == -1 {
			fmt.Printf("∞\t%d\n", daysSinceLastReviewed[days])
		} else {
			fmt.Printf("%d\t%d\n", days, daysSinceLastReviewed[days])
		}
	}
}
