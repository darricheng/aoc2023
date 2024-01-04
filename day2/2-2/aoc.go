package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Reasoning
*/

func split(r rune) bool {
	return r == ',' || r == ';'
}

func main() {
	b, err := os.ReadFile("../data.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	data := strings.Split(input, "\n")
	numOfLines := len(data)

	sum := 0

	for i, line := range data {
		if i == numOfLines-1 {
			break
		}
		fmt.Println("line", i, "--", line)

		// Code goes here
		tmp := strings.Split(line, ":")

		picks := strings.FieldsFunc(tmp[1], split)
		red := 0
		green := 0
		blue := 0
		for _, pick := range picks {
			x := strings.Split(strings.TrimSpace(pick), " ")
			color := x[1]
			qty, err := strconv.Atoi(x[0])
			if err != nil {
				panic(err)
			}
			switch color {
			case "red":
				if qty > red {
					red = qty
				}
			case "green":
				if qty > green {
					green = qty
				}
			case "blue":
				if qty > blue {
					blue = qty
				}
			}
		}
		println(red, green, blue)
		sum += red * green * blue
	}

	fmt.Println("Final Result:", sum)
}
