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

	res := 0

	for i, line := range data {
		if i == numOfLines-1 {
			break
		}
		fmt.Println("line", i, "--", line)

		// Code goes here
		strSlice := strings.Split(line, " ")
		var numSlice []int
		for _, str := range strSlice {
			num, ok := strconv.Atoi(str)
			if ok != nil {
				panic("num parse failed")
			}
			numSlice = append(numSlice, num)
		}
		res += recursive(numSlice)
	}

	fmt.Printf("FINAL RESULT: %d\n", res)
}

func recursive(slice []int) int {
	allZero := true
	if slice[0] != 0 {
		allZero = false
	}
	var differences []int
	for i := 1; i < len(slice); i++ {
		if allZero && slice[i] != 0 {
			allZero = false
		}
		differences = append(differences, slice[i]-slice[i-1])
	}
	if allZero {
		return 0
	} else {
		return slice[len(slice)-1] + recursive(differences)
	}
}
