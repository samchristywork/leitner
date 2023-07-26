package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type historyLine struct {
	card          string
	correct       bool
	last_reviewed int64
}

func readHistory(filename string) []historyLine {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading " + filename + ". You should create it if it doesn't exist.")
		os.Exit(1)
	}

	lines := make([]historyLine, 0)

	for _, line := range strings.Split(string(contents), "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "	")

		if parts[0] == "quiz" && len(parts) == 3 {
			lines = append(lines, historyLine{parts[1], parts[2] == "correct", 0})
			continue
		} else if parts[0] == "quiz" && len(parts) == 4 {
			timestamp, err := strconv.Atoi(parts[3])
			if err == nil {
				lines = append(lines, historyLine{parts[1], parts[2] == "correct", int64(timestamp)})
				continue
			}
		}

		fmt.Println("Error reading " + filename + ".")
		os.Exit(1)
	}

	return lines
}
