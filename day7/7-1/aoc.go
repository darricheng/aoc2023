package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
Use the slices.SortFunc to sort the hands with the comparator function.
Hands with fewer different cards have higher rank.
Five of a kind has 1 card, while four of a kind and full house have 2 cards.
Then compare within the number of cards. Then compare by high card.
The comparator function will then return positive or negative as per the comparison.
Then get the score with the rank. For the sorting, whole thing is a slice of arrays of length 2, where [hand, score].
*/

func main() {
	b, err := os.ReadFile("../data.txt")
	if err != nil {
		panic(err)
	}
	input := string(b)
	data := strings.Split(input, "\n")
	numOfLines := len(data)

	res := 0

	var hands [][]string

	for i, line := range data {
		if i == numOfLines-1 {
			break
		}
		fmt.Println("line", i, "--", line)

		// Code goes here
		hands = append(hands, strings.Split(line, " "))
	}

	// sort the hands
	slices.SortFunc(hands, sortAlgo)
	// fmt.Println(hands)

	for i, hand := range hands {
		score, err := strconv.Atoi(hand[1])
		if err != nil {
			panic(err)
		}
		res += (i + 1) * score
	}

	// Total up the scores

	fmt.Printf("FINAL RESULT: %d\n", res)
}

var cardStr = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

// NOTE: Return negative if a < b
func sortAlgo(a, b []string) int {
	freqCountA := make(map[rune]int)
	freqCountB := make(map[rune]int)

	for _, char := range a[0] {
		_, ok := freqCountA[char]
		if ok {
			freqCountA[char]++
		} else {
			freqCountA[char] = 1
		}
	}
	for _, char := range b[0] {
		_, ok := freqCountB[char]
		if ok {
			freqCountB[char]++
		} else {
			freqCountB[char] = 1
		}
	}

	numOfDiffCardsA := len(freqCountA)
	numOfDiffCardsB := len(freqCountB)

	// More cards means lower rank
	if numOfDiffCardsA > numOfDiffCardsB {
		return -1
	}
	if numOfDiffCardsA < numOfDiffCardsB {
		return 1
	}

	// Both hands have the same number of different cards
	// Here, we check the different types of hands
	var comparator int
	switch numOfDiffCardsA {
	case 1: // do nothing
	case 2:
		comparator = compareSameNumOfDiffCards(freqCountA, freqCountB)
	case 3:
		comparator = compareSameNumOfDiffCards(freqCountA, freqCountB)
	case 4: // do nothing
	case 5: // do nothing
	}
	if comparator != 0 {
		return comparator
	}

	// Check high card
	var highCardCheck int
	handA := []rune(a[0])
	handB := []rune(b[0])
	for i := 0; i < len(handA); i++ {
		if handA[i] != handB[i] {
			indexA := slices.Index(cardStr, handA[i])
			indexB := slices.Index(cardStr, handB[i])
			highCardCheck = indexB - indexA
			break
		}
	}

	return highCardCheck
}

func compareSameNumOfDiffCards(mapA, mapB map[rune]int) int {
	var mostA int
	var mostB int
	for _, freq := range mapA {
		if freq > mostA {
			mostA = freq
		}
	}
	for _, freq := range mapB {
		if freq > mostB {
			mostB = freq
		}
	}
	return mostA - mostB
}
