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

	/*
		To find the area, we go through every single grid one by one left to right, top to bottom
		with a state that indicates whether that grid is enclosed by the path or not.
		The simplest is crossing a |, where we will just toggle on or off depending on the current state.
		Crossing a - does nothing, we just continue iterating through.
		Crossing a corner is more complicated.

		{
			x: {
				y: "identifier"
			}
		}

		WARN: The below algo doesn't work. Simple way to think of it is J at bottom right of a square path shouldn't toggle on.
		Maybe one way is to accumulate the value when toggled on, but only add to final when toggled off?
		But I guess it doesn't work either. Consider the following shape:
		F-------7
		|.F---7.|
		|.|...|.|
		L-J...L-J
		The bottom row would toggle on at the bottom left J and off at the next L, resulting in three grids falsely added.

		identifier is any of the following:
		- "B": toggles on and off
		- "Y": toggles on only
		- "N": toggles off only
		- "O": does nothing

		J and 7 toggle on only
		F and L toggle off only
		| toggles on and off
		- does nothing

		S in the puzzle input is an F, so it should toggle off only
		but it might differ for other inputs
	*/
	var h map[int]map[int]string

	// trace path and build map of toggles
	for {
		c := m[y][x]
		if c == "S" {
			// NOTE: needs to be changed according to the puzzle input
			h[x][y] = "N"
			break
		}
		// create hashmap of toggles
		switch c {
		case "J":
			h[x][y] = "Y"
		case "7":
			h[x][y] = "Y"
		case "F":
			h[x][y] = "N"
		case "L":
			h[x][y] = "N"
		case "|":
			h[x][y] = "B"
		case "-":
			h[x][y] = "O"
		}
		// trace path
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

	// go through grid and calculate area
	res := 0
	for y, row := range m {
		toggle := false
		accumulator := 0
		for x := range row {
			if toggle {
				accumulator++
			}
			xMap, xOk := h[x]
			if !xOk {
				continue
			}
			toggleInstruction, yOk := xMap[y]
			if !yOk {
				continue
			}
			switch toggleInstruction {
			case "B":
				toggle = !toggle
				if !toggle {
					accumulator--
					res += accumulator
					accumulator = 0
				}
			case "Y":
				toggle = true
			case "N":
				toggle = false
				accumulator--
				res += accumulator
				accumulator = 0
			case "O":
			}
		}
	}

	fmt.Printf("FINAL RESULT: %d\n", res)
}
