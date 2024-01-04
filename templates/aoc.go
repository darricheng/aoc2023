package main

import (
	"fmt"
	"os"
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

	for i, line := range data {
		if i == numOfLines-1 {
			break
		}
		fmt.Println("line", i, "--", line)

		// Code goes here

	}
}
