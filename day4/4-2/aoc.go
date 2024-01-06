package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
Algorithm

1. Increment the final res by the number of current cards
2. Find the number of winning numbers for the current card
3. Increment the count of the subsequent cards accordingly by the number of current cards
*/

func main() {
	b, err := os.ReadFile("../data.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	data := strings.Split(input, "\n")
	numOfLines := len(data)

	res := 0

	// init map of card quantities
	cardQtys := make(map[int]int)
	for i := 0; i < len(data)-1; i++ {
		cardQtys[i+1] = 1
	}

	for i, line := range data {
		if i == numOfLines-1 {
			break
		}
		fmt.Println("line", i, "--", line)

		// Code goes here
		currentCardNum := i + 1

		// increment the final res
		res += cardQtys[currentCardNum]

		// Count number of winning numbers for the current card
		tmp := strings.Split(line, ":")
		nums := strings.Split(strings.TrimSpace(tmp[1]), "|")
		winningNums := strings.Split(strings.TrimSpace(nums[0]), " ")
		elfNums := strings.Split(strings.TrimSpace(nums[1]), " ")

		winningNums = whiteSpaceBegone(winningNums)
		elfNums = whiteSpaceBegone(elfNums)

		wins := 0
		for _, n := range elfNums {
			if slices.Contains(winningNums, n) {
				wins++
			}
		}

		// increment totals of subsequent cards accordingly
		for i := currentCardNum + 1; i <= currentCardNum+wins; i++ {
			cardQtys[i] += cardQtys[currentCardNum]
		}
	}

	fmt.Printf("FINAL RESULT: %d", res)
}

func whiteSpaceBegone(strSlice []string) []string {
	newStrSlice := []string{}
	for _, char := range strSlice {
		if _, err := strconv.Atoi(char); err == nil {
			newStrSlice = append(newStrSlice, char)
		}
	}
	return newStrSlice
}
