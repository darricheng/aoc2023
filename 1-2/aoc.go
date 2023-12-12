package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// b, err := os.ReadFile("input.txt")
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// input := string(b)
	// data := strings.Split(input, "\n")

	data := make([]string, 1)
	data[0] = "one4seveneight"

	// Code goes here
	sum := 0

	for _, line := range data {
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
		fmt.Println("beforeFirstDigitSlice: ", beforeFirstDigitSlice)

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
		fmt.Println("firstNum: ", firstNum)
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
		fmt.Println("firstDigit: ", firstDigit)

		if lastDigitIndex == -1 {
			lastDigitIndex = 0
		}
		afterLastDigitSlice := line[lastDigitIndex+1:]
		fmt.Println("afterLastDigitSlice: ", afterLastDigitSlice)

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
		fmt.Println("lastNum: ", lastNum)
		if lastNum != -1 {
			for _, numStr := range strconv.Itoa(lastNum) {
				lastDigit = numStr
			}
		}
		fmt.Println("lastDigit: ", lastDigit)

		var sb strings.Builder
		sb.WriteRune(firstDigit)
		sb.WriteRune(lastDigit)
		fmt.Println("sb: ", sb)
		str := sb.String()
		fmt.Println("str: ", str)
		i, err := strconv.Atoi(str)
		if err != nil {
		}
		fmt.Println("i: ", i)
		sum += i
	}

	fmt.Println(sum)
}
