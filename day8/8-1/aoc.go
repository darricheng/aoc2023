package main

import (
	"fmt"
	"os"
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

	instructions := strings.Split(data[0], "")

	fullMap := make(map[string]map[string]string)

	for i, line := range data {
		if i == numOfLines-1 {
			break
		}
		if i == 1 || i == 0 {
			continue
		}
		// fmt.Println("line", i, "--", line)

		// Code goes here
		key, vals, _ := strings.Cut(line, " = ")
		noParenVals := strings.Trim(vals, "()")
		left, right, _ := strings.Cut(noParenVals, ", ")
		fullMap[key] = map[string]string{"L": left, "R": right}
	}

	current := "AAA"

main:
	for {
		for _, str := range instructions {
			current = fullMap[current][str]
			res++
			if current == "ZZZ" {
				break main
			}
		}
	}

	fmt.Printf("FINAL RESULT: %d\n", res)
}
