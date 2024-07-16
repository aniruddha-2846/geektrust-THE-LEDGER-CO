package main

import (
	"bufio"
	"fmt"
	"geektrust/subpackages"
	"os"
	"strings"
)

func getFirstWord(s string) string {
	words := strings.Fields(s)
	if len(words) > 0 {
		return words[0]
	}
	return ""
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please give the input file path")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening the file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		firstWord := getFirstWord(line)

		switch firstWord {
		case "LOAN":
			subpackages.ProcessLoan(line)
		case "PAYMENT":
			subpackages.ProcessPayment(line)
		case "BALANCE":
			subpackages.ProcessBalance(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading the file")
		os.Exit(1)
	}
}
