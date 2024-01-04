package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
Algorithm
1. Iterate through the string to find the first number
2. Extend until found the complete number
3. Use the index and length of the number to search around the number for symbols
4. If symbols are found, add the number to the res
*/

func findFullNum(str string, n int) []int {
	runes := []rune(str)
	res := string(runes[n])
	for i := n + 1; i < len(runes); i++ {
		if !unicode.IsDigit(runes[i]) {
			break
		}
		res = strings.Join([]string{res, string(runes[i])}, "")
	}
	finalRes, err := strconv.Atoi(res)
	if err != nil {
		panic(err)
	}
	return []int{finalRes, len(res)}
}

func subchecker(row int, col int, data []string) bool {
	x := rune(data[row][col])
	return !unicode.IsDigit(x) && x != '.'
}

func valid(num int, length int, row int, index int, data []string) bool {
	println("checking:", num)
	startIndex := 0
	endIndex := len(data[row]) - 1
	// char before
	if index != startIndex {
		startIndex = index - 1
		if subchecker(row, startIndex, data) {
			return true
		}
	}
	// char after
	if index+length <= endIndex {
		endIndex = index + length
		if subchecker(row, index+length, data) {
			return true
		}
	}
	// line above
	if row != 0 {
		for i := startIndex; i <= endIndex; i++ {
			if subchecker(row-1, i, data) {
				return true
			}
		}
	}
	// line below
	if row != len(data)-2 {
		for i := startIndex; i <= endIndex; i++ {
			if subchecker(row+1, i, data) {
				return true
			}
		}
	}

	return false
}

func main() {
	b, err := os.ReadFile("../data.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	data := strings.Split(input, "\n")
	numOfLines := len(data)

	res := 0

	for i, line := range data {
		if i == numOfLines-1 {
			break
		}
		fmt.Println("line", i, "--", line)

		counter := 0 // for skipping chars in the loop when a num is found
		// Code goes here
		for j, char := range line {
			if counter != 0 {
				counter--
				continue
			}
			if unicode.IsDigit(char) {
				tmp := findFullNum(line, j)
				currentNum := tmp[0]
				numLen := tmp[1]
				counter += numLen

				if valid(currentNum, numLen, i, j, data) {
					println("valid!")
					res += currentNum
				}
			}
		}
	}
	fmt.Println("FINAL RESULT:", res)
}
