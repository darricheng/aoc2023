package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Reasoning
Iterate through each color and number, check if any is more than the 12/13/14 of each color.
I think it shouldn't matter to check by groups, since the only way to determine if a group
is valid is whether there are any colors that exceed the 12/13/14.
*/

func split(r rune) bool {
	return r == ',' || r == ';'
}
func possible(s string, q int) bool {
	var res bool
	switch s {
	case "red":
		res = q <= 12
	case "green":
		res = q <= 13
	case "blue":
		res = q <= 14
	}
	return res
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
		game := tmp[0]

		gameNum, err := strconv.Atoi(strings.TrimPrefix(game, "Game "))
		if err != nil {
			panic(err)
		}
		println(gameNum)

		picks := strings.FieldsFunc(tmp[1], split)
		valid := true
		for _, pick := range picks {
			x := strings.Split(strings.TrimSpace(pick), " ")
			color := x[1]
			qty, err := strconv.Atoi(x[0])
			if err != nil {
				panic(err)
			}
			if !possible(color, qty) {
				valid = false
				break
			}
		}
		if valid {
			sum += gameNum
		}
	}

	fmt.Println("Final Result:", sum)
}
