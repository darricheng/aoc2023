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
1. Identify row and col numbers for the asterisks
2. For each asterisk, perform the subsequent steps
3. Create a slice to store valid numbers
4. If the col before is a number, extend back to get the number and push it into slice
5. If the col after is a number, extend front to get the number and push it into slice
6. Check the above row if the current row isn't 0
6a. If the col before is a number, extend back to get the number and push it into slice
6b. Search forwards, ending the search at the col after index, push valid nums into slice
7. Repeat step 6 for the row below if it's not the second last row (last row will be empty)
8. If len(slice) == 2, multiply the numbers in the slice and add it to the total sum
*/

var data []string

func main() {
	b, err := os.ReadFile("../data.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	data = strings.Split(input, "\n")
	numOfLines := len(data)

	res := 0

	for row, line := range data {
		if row == numOfLines-1 {
			break
		}
		fmt.Println("line", row, "--", line)

		// Code goes here
		for col, char := range line {
			if char == '*' {
				adjNums := []int{}
				firstCol := col == 0
				lastCol := col == len(data[row])-1
				// backwards
				if !firstCol && isNum(row, col-1) {
					adjNums = append(adjNums, getBackwardsNum(row, col-1))
				}
				// forwards
				if !lastCol && isNum(row, col+1) {
					adjNums = append(adjNums, getForwardsNum(row, col+1))
				}
				// row above
				if row != 0 {
					adjNums = append(adjNums, getAdjRowNum(row-1, col)...)

				}
				// row below
				if row != len(data)-2 {
					adjNums = append(adjNums, getAdjRowNum(row+1, col)...)
				}

				if len(adjNums) == 2 {
					fmt.Println(adjNums)
					res += adjNums[0] * adjNums[1]
				}
			}
		}
	}
	fmt.Println("FINAL RESULT:", res)
}

func isNum(row int, col int) bool {
	_, err := strconv.Atoi(string(rune(data[row][col])))
	return err == nil
}

func getAdjRowNum(row int, col int) []int {
	res := []int{}
	/*
		3 possible positions for *
		1. col 0
		2. last col
		3. somewhere in between

		For 1 and 2, if current col isDigit, work from there, otherwise check diagonal and start there

		For 3, I want to try using bitflags to indicate the different scenarios
		- 000: do nothing
		- 010: single number
		- 101: search backwards and forwards respectively
		- 001 & 011: search forwards
		- 100 & 110: search backwards
		- 111: search backwards then forwards
	*/
	if col == 0 { // col 0
		if isNum(row, col) { // 1x
			res = append(res, getForwardsNum(row, col))
		} else if isNum(row, col+1) { // 01
			res = append(res, getForwardsNum(row, col+1))
		}
	} else if col == len(data[row])-1 { // last col
		if isNum(row, col) { // x1
			res = append(res, getBackwardsNum(row, col))
		} else if isNum(row, col-1) { // 10
			res = append(res, getBackwardsNum(row, col-1))
		}
	} else { // somewhere in between
		flags := 0
		if isNum(row, col-1) {
			flags = flags | 0b_100
		}
		if isNum(row, col) {
			flags = flags | 0b_010
		}
		if isNum(row, col+1) {
			flags = flags | 0b_001
		}
		// Print the flags in binary format
		// fmt.Printf("%03b\n", flags)
		switch flags {
		case 0b_000:
			// do nothing
		case 0b_010:
			tmp, err := strconv.Atoi(string(rune(data[row][col])))
			if err != nil {
				panic(err)
			}
			res = append(res, tmp)
		case 0b_101:
			res = append(res, getBackwardsNum(row, col-1))
			res = append(res, getForwardsNum(row, col+1))
		// forwards
		case 0b_001:
			res = append(res, getForwardsNum(row, col+1))
		case 0b_011:
			res = append(res, getForwardsNum(row, col))
		// backwards
		case 0b_100:
			res = append(res, getBackwardsNum(row, col-1))
		case 0b_110:
			res = append(res, getBackwardsNum(row, col))
		// 111
		case 0b_111:
			backPart := strconv.Itoa(getBackwardsNum(row, col))
			/*
				The below line only works because there are no numbers larger than 3 digits
				If there are numbers found by the below function that starts with 0, the returned result will be less a 0 in the middle of the number
				A fix could involve changing the getForwardsNum to return a string, then parse the string in the calling function
			*/
			frontPart := strconv.Itoa(getForwardsNum(row, col+1))
			fullNum, err := strconv.Atoi(backPart + frontPart)
			if err != nil {
				panic(err)
			}
			res = append(res, fullNum)
		}
	}

	return res
}

func getForwardsNum(row int, col int) int {
	runes := []rune(data[row])
	res := string(runes[col])
	for i := col + 1; i < len(runes); i++ {
		if !unicode.IsDigit(runes[i]) {
			break
		}
		res = strings.Join([]string{res, string(runes[i])}, "")
	}
	finalRes, err := strconv.Atoi(res)
	if err != nil {
		panic(err)
	}
	return finalRes
}

func getBackwardsNum(row int, col int) int {
	runes := []rune(data[row])
	res := string(runes[col])
	for i := col - 1; i >= 0; i-- {
		if !unicode.IsDigit(runes[i]) {
			break
		}
		res = strings.Join([]string{string(runes[i]), res}, "")
	}
	finalRes, err := strconv.Atoi(res)
	if err != nil {
		panic(err)
	}
	return finalRes
}
