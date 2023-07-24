package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	deckdir     string
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
		fmt.Println("Config file (" + filename + ") not found")
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
