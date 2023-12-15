package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	data := strings.Split(input, "\n")

	// data := make([]string, 5)
	// data[0] = "one4seveneight"
	// data[1] = "ne4eveneght"
	// data[2] = "oneseveneit"
	// data[3] = "one4sevenei9ght"
	// data[4] = "two4seveneight5"

	// Code goes here
	sum := 0

	for _, line := range data {
		// break at the last data point
		if len(line) == 0 {
			break
		}
		firstDigit := '0'
		lastDigit := '0'
		firstDigitIndex := -1
		lastDigitIndex := -1

		for i, char := range line {
			if unicode.IsDigit(char) {
				firstDigit = char
				firstDigitIndex = i
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			char := rune(line[i])
			if unicode.IsDigit(char) {
				lastDigit = char
				lastDigitIndex = i
				break
			}
		}

		n := make([]string, 10)
		n[0] = "zero"
		n[1] = "one"
		n[2] = "two"
		n[3] = "three"
		n[4] = "four"
		n[5] = "five"
		n[6] = "six"
		n[7] = "seven"
		n[8] = "eight"
		n[9] = "nine"

		if firstDigitIndex == -1 {
			firstDigitIndex = len(line) - 1
		}
		beforeFirstDigitSlice := line[:firstDigitIndex]

		smallestFirst := 10
		firstNum := -1
		for i, s := range n {
			if strings.Contains(beforeFirstDigitSlice, s) {
				index := strings.Index(beforeFirstDigitSlice, s)
				if index < smallestFirst {
					smallestFirst = index
					firstNum = i
				}
			}
		}
		if firstNum != -1 {
			// NOTE: rune is an alias for int32
			// Each rune can only represent a single byte
			// So we need to loop through the string to extract the individual runes
			// But since we're sure that firstNum will always be a string of length 1,
			// we can just extract the rune with a loop as per below
			for _, numStr := range strconv.Itoa(firstNum) {
				firstDigit = numStr
			}
		}

		if lastDigitIndex == -1 {
			lastDigitIndex = 0
		}
		afterLastDigitSlice := line[lastDigitIndex+1:]

		smallestLast := lastDigitIndex + 1
		lastNum := -1
		for i, s := range n {
			if strings.Contains(afterLastDigitSlice, s) {
				index := strings.Index(afterLastDigitSlice, s)
				if index > smallestLast {
					smallestLast = index
					lastNum = i
				}
			}
		}
		if lastNum != -1 {
			for _, numStr := range strconv.Itoa(lastNum) {
				lastDigit = numStr
			}
		}

		var sb strings.Builder
		sb.WriteRune(firstDigit)
		sb.WriteRune(lastDigit)
		str := sb.String()
		i, err := strconv.Atoi(str)
		if err != nil {
		}
		sum += i
		// fmt.Println(dataI, ":", sum)
	}

	fmt.Println(sum)
}
