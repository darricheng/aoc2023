package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

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

		// Code goes here
		tmp := strings.Split(line, ":")
		nums := strings.Split(strings.TrimSpace(tmp[1]), "|")
		winningNums := strings.Split(strings.TrimSpace(nums[0]), " ")
		elfNums := strings.Split(strings.TrimSpace(nums[1]), " ")

		winningNums = whiteSpaceBegone(winningNums)
		elfNums = whiteSpaceBegone(elfNums)

		counter := 0
		for _, n := range elfNums {
			if slices.Contains(winningNums, n) {
				counter++
			}
		}
		switch counter {
		case 0:
		default:
			inc := int(math.Pow(2.0, float64(counter-1)))
			res += inc
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
