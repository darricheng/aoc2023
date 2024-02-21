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

	routes := []string{}

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

		if key[2] == 'A' {
			routes = append(routes, key)
		}
	}

	var multiples []int

	for _, start := range routes {
		timesToZ := 0
		current := start
	main:
		for {
			for _, str := range instructions {
				timesToZ++
				current = fullMap[current][str]
				if current[2] == 'Z' {
					break main
				}
			}
		}
		multiples = append(multiples, timesToZ)
	}
	res = LCM(multiples[0], multiples[1], multiples...)

	// So brute forcing doesn't seem to be the way.
	// I initially thought of lowest common multiple, but I put it aside because I'm not sure whether
	// a loop even exists for all the starting values, i.e. each starting point lands on a Z every X times.
	// 	routesLen := len(routes)
	// 	println("routesLen:", routesLen)
	// main:
	// 	for {
	// 		for _, str := range instructions {
	// 			numOfZs := 0
	// 			res++
	// 			msg := make(chan int)
	// 			for i, current := range routes {
	// 				go func(i int, current string) {
	// 					routes[i] = fullMap[current][str]
	// 					if routes[i][2] == 'Z' {
	// 						msg <- 1
	// 					} else {
	// 						msg <- 0
	// 					}
	// 				}(i, current)
	// 			}
	// 			for i := 0; i < routesLen; i++ {
	// 				rcv := <-msg
	// 				numOfZs += rcv
	// 			}
	// 			if numOfZs == len(routes) {
	// 				break main
	// 			}
	// 		}
	// 	}

	fmt.Printf("FINAL RESULT: %d\n", res)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
