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

	// Code goes here
	sum := 0

	for _, line := range data {
		firstDigit := '0'
		lastDigit := '0'
		for _, char := range line {
			if unicode.IsDigit(char) {
				if firstDigit == '0' {
					firstDigit = char
				}
				lastDigit = char
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
	}

	fmt.Println(sum)
}
