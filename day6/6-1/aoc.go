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

	var nums [][]int

	res := 1
	for i, line := range data {
		if i == numOfLines-1 {
			break
		}
		fmt.Println("line", i, "--", line)

		// Code goes here
		tmpData := strings.Split(line, ":")
		tmpNums := strings.Split(strings.TrimSpace(tmpData[1]), " ")
		nums = append(nums, whiteSpaceBegoneAndConvertToInt(tmpNums))
	}

	times := nums[0]
	distances := nums[1]

	// Iterate through the times
	for i, time := range times {
		dist := distances[i]
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
	}

	fmt.Printf("FINAL RESULT: %d\n", res)
}

func whiteSpaceBegoneAndConvertToInt(strSlice []string) []int {
	newIntSlice := []int{}
	for _, char := range strSlice {
		if i, err := strconv.Atoi(char); err == nil {
			newIntSlice = append(newIntSlice, i)
		}
	}
	return newIntSlice
}
