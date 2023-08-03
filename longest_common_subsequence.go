package main

import (
	"fmt"
)

const (
	unchanged = 0
	added     = 1
	removed   = 2
)

type token struct {
	kind int
	val  interface{}
}

func red_string() string {
	return "\033[31m"
}

func green_string() string {
	return "\033[32m"
}

func normal_string() string {
	return "\033[0m"
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(r []interface{}, c []interface{}) [][]int {
	table := make([][]int, len(r)+1)
	for i := range table {
		table[i] = make([]int, len(c)+1)
	}

	for i := 1; i < len(r)+1; i++ {
		for j := 1; j < len(c)+1; j++ {
			if r[i-1] == c[j-1] {
				table[i][j] = table[i-1][j-1] + 1
			} else {
				table[i][j] = max(table[i-1][j], table[i][j-1])
			}
		}
	}

	return table
}

func diff(r []interface{}, c []interface{}, table [][]int) []token {
	var output []token

	i := len(r)
	j := len(c)
	for i > 0 && j > 0 {
		if r[i-1] == c[j-1] {
			output = append(output, token{unchanged, r[i-1]})
			i--
			j--
		} else if table[i-1][j] > table[i][j-1] {
			output = append(output, token{removed, r[i-1]})
			i--
		} else {
			output = append(output, token{added, c[j-1]})
			j--
		}
	}

	for i > 0 {
		output = append(output, token{removed, r[i-1]})
		i--
	}

	for j > 0 {
		output = append(output, token{added, c[j-1]})
		j--
	}

	return output
}

func printDiff(output []token) {
	for i := len(output) - 1; i >= 0; i-- {
		switch output[i].kind {
		case unchanged:
			fmt.Print(normal_string())
		case added:
			fmt.Print(green_string())
		case removed:
			fmt.Print(red_string())
		}
		fmt.Print(output[i].val)

		if i > 0 {
			fmt.Print(" ")
		} else {
			fmt.Print(normal_string())
		}
	}

	fmt.Println()
}

func tokenizeString(s string) []interface{} {
	if s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}

	var output []interface{}
	var current string

	for _, c := range s {
		if c == ' ' {
			output = append(output, current)
			current = ""
		} else {
			current += string(c)
		}
	}

	if current != "" {
		output = append(output, current)
	}

	return output
}

func tokenize[T any](s []T) []interface{} {
	var output []interface{}
	for _, p := range s {
		output = append(output, p)
	}

	return output
}
