package main

import (
	"fmt"
	"os"
	"strings"
)

type Flashcard struct {
	deck          string
	front         string
	back          string
	bin           uint32
	last_reviewed int64
}

func (f Flashcard) String() string {
	return fmt.Sprintf("%s|%s|%s", f.deck, f.front, f.back)
}

func (f Flashcard) hashFlashcard() uint32 {
	return hashString(f.String())
}

func parseFlashcard(s string) Flashcard {
	segments := strings.Split(s, "|")
	if len(segments) != 4 {
		fmt.Printf("Error: flashcard has wrong number of segments: %s\n", s)
		os.Exit(1)
	}

	deck := segments[1]
	front := segments[2]
	back := segments[3]

	return Flashcard{deck, front, back, 1, 0}
}
