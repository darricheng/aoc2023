package main

import (
	"fmt"
	"os"
	"slices"
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

	// accessing m is done with m[y][x]
	var m [][]string
	x := -1
	y := -1

	for i, line := range data {
		if i == numOfLines-1 {
			break
		}
		// Code goes here
		tmp := strings.Split(line, "")
		sIndex := slices.Index(tmp, "S")
		if sIndex >= 0 {
			x = sIndex
			y = i
		}
		m = append(m, tmp)
	}

	count := 1
	dir := ""

	startFound := false
	tmp := ""
	// find a start
	tmp = m[y][x+1]
	if tmp == "7" || tmp == "J" || tmp == "-" {
		startFound = true
		x = x + 1
		dir = "R"
	}
	tmp = m[y][x-1]
	if !startFound && (tmp == "L" || tmp == "F" || tmp == "-") {
		startFound = true
		x = x - 1
		dir = "L"
	}
	tmp = m[y+1][x]
	if !startFound && (tmp == "L" || tmp == "J" || tmp == "|") {
		startFound = true
		y = y + 1
		dir = "D"
	}
	tmp = m[y-1][x]
	if !startFound && (tmp == "7" || tmp == "F" || tmp == "|") {
		startFound = true
		y = y - 1
		dir = "U"
	}

	for {
		count++
		c := m[y][x]
		println(c)
		if c == "S" {
			break
		}
		switch dir {
		case "R":
			switch c {
			case "-":
				x++
			case "J":
				y--
				dir = "U"
			case "7":
				y++
				dir = "D"
			}
		case "L":
			switch c {
			case "-":
				x--
			case "F":
				y++
				dir = "D"
			case "L":
				y--
				dir = "U"
			}
		case "U":
			switch c {
			case "|":
				y--
			case "F":
				x++
				dir = "R"
			case "7":
				x--
				dir = "L"
			}
		case "D":
			switch c {
			case "|":
				y++
			case "J":
				x--
				dir = "L"
			case "L":
				x++
				dir = "R"
			}
		}
	}

	fmt.Printf("FINAL RESULT: %d\n", count/2)
}
