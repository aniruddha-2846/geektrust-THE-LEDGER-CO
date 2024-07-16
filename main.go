package main

import (
	"bufio"
	"geektrust/subpackages"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err1 := os.Open(filename)
	if err1 != nil {
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		firstWord := subpackages.GetFirstWord(line)
		switch firstWord {
		case "LOAN":
			subpackages.ProcessLoan(line)
		case "PAYMENT":
			subpackages.ProcessPayment(line)
		case "BALANCE":
			subpackages.ProcessBalance(line)
		}
	}
}
