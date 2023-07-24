package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type historyLine struct {
	card    string
	correct bool
}

func readHistory(filename string) []historyLine {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading "+filename+". You should create it if it doesn't exist.")
		os.Exit(1)
	}

	lines := make([]historyLine, 0)

	for _, line := range strings.Split(string(contents), "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "	")
		if len(parts) != 2 {
			fmt.Println("Error reading "+filename+".")
			os.Exit(1)
		}

		lines = append(lines, historyLine{parts[0], parts[1] == "correct"})
	}

	return lines
}
