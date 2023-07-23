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

func readHistory() []historyLine {
	contents, err := ioutil.ReadFile("/home/sam/.flash_history")
	if err != nil {
		fmt.Println("Error reading ~/.flash_history. You should create it if it doesn't exist.")
		os.Exit(1)
	}

	lines := make([]historyLine, 0)

	for _, line := range strings.Split(string(contents), "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "	")
		if len(parts) != 2 {
			fmt.Println("Error reading ~/.flash_history.")
			os.Exit(1)
		}

		lines = append(lines, historyLine{parts[0], parts[1] == "correct"})
	}

	return lines
}