package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("../data.txt")
	if err != nil {
		panic(err)
	}
	input := string(b)
	data := strings.Split(input, "\n")
	numOfLines := len(data)

	var nums [][]string

	res := 1
	for i, line := range data {
		if i == numOfLines-1 {
			break
		}
		fmt.Println("line", i, "--", line)

		// Code goes here
		tmpData := strings.Split(line, ":")
		tmpNums := strings.Split(strings.TrimSpace(tmpData[1]), " ")
		nums = append(nums, whiteSpaceBegone(tmpNums))
	}

	time, err := strconv.Atoi(strings.Join(nums[0], ""))
	if err != nil {
		panic(err)
	}
	dist, err := strconv.Atoi(strings.Join(nums[1], ""))
	if err != nil {
		panic(err)
	}
	// Iterate through the times
	timeIsEven := time%2 == 0

	count := 0

	for j := time / 2; j*(time-j) > dist; j-- {
		count++
	}

	count *= 2
	if timeIsEven {
		count--
	}

	res *= count

	fmt.Printf("FINAL RESULT: %d\n", res)
}

func whiteSpaceBegone(strSlice []string) []string {
	newIntSlice := []string{}
	for _, char := range strSlice {
		if _, err := strconv.Atoi(char); err == nil {
			newIntSlice = append(newIntSlice, char)
		}
	}
	return newIntSlice
}
