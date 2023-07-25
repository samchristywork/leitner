package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	deck_dir     string
	suffix string
	question_filename string
	history_filename string
	incorrect_review_time int
}

func dumpConfig(config Config) {
	fmt.Println("deck_dir: " + config.deck_dir)
	fmt.Println("suffix: " + config.suffix)
	fmt.Println("question_filename: " + config.question_filename)
	fmt.Println("history_filename: " + config.history_filename)
	fmt.Println("incorrect_review_time: " + strconv.Itoa(config.incorrect_review_time))
}

func readConfigFile(filename string) Config {
	config := Config{}

	home := os.Getenv("HOME")

	// Defaults
	config.deck_dir = "./"
	config.suffix = ".txt"
	config.question_filename = "/tmp/flash_question"
	config.history_filename = home + "/.flash_history"
	config.incorrect_review_time = 10

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
		if strings.HasPrefix(line, "deck_dir:") {
			config.deck_dir = strings.TrimSpace(strings.TrimPrefix(line, "deck_dir:"))
		}
		if strings.HasPrefix(line, "suffix:") {
			config.suffix = strings.TrimSpace(strings.TrimPrefix(line, "suffix:"))
		}
		if strings.HasPrefix(line, "question_filename:") {
			config.question_filename = strings.TrimSpace(strings.TrimPrefix(line, "question_filename:"))
		}
		if strings.HasPrefix(line, "history_filename:") {
			config.history_filename = strings.TrimSpace(strings.TrimPrefix(line, "history_filename:"))
		}
		if strings.HasPrefix(line, "incorrect_review_time:") {
			incorrect_review_time, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "incorrect_review_time:")))
			if err != nil {
				fmt.Println("Error reading incorrect_review_time")
				os.Exit(0)
			}
			config.incorrect_review_time = incorrect_review_time
		}
	}

	return config
}
