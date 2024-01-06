package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type mapNums struct {
	destStart   int
	sourceStart int
	rangeLen    int
}

var seeds []int
var seedToSoil []mapNums
var soilToFertilizer []mapNums
var fertilizerToWater []mapNums
var waterToLight []mapNums
var lightToTemp []mapNums
var tempToHumidity []mapNums
var humidityToLocation []mapNums

func main() {
	b, err := os.ReadFile("../data.txt")
	if err != nil {
		panic(err)
	}
	input := string(b)
	data := strings.Split(input, "\n")
	numOfLines := len(data)

	res := math.MaxInt
	categoryCounter := 0
	intraCounter := 0
	var currentMap *[]mapNums

	// parse data into respective vars
	for i, line := range data {
		if i == numOfLines-1 {
			break
		}
		// fmt.Println("line", i, "--", line)

		switch categoryCounter {
		case 0: // seeds
			seedNums := strings.Split(line, " ")[1:] // discard first val which is "seeds:"
			for _, s := range seedNums {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, n)
			}
		case 1: // seed-to-soil
			currentMap = &seedToSoil
		case 2: // soil-to-fertilizer
			currentMap = &soilToFertilizer
		case 3: // fertilizer-to-water
			currentMap = &fertilizerToWater
		case 4: // water-to-light
			currentMap = &waterToLight
		case 5: // light-to-temperature
			currentMap = &lightToTemp
		case 6: // temperature-to-humidity
			currentMap = &tempToHumidity
		case 7: // humidity-to-location
			currentMap = &humidityToLocation
		}

		// empty line means we are parsing the next category
		if len(line) == 0 {
			categoryCounter++
			intraCounter = 0
		} else {
			intraCounter++
		}

		// don't continue executing if processing seeds
		if categoryCounter == 0 {
			continue
		}
		// parse data into map
		// intraCounter is 0 for blank line and 1 for heading
		if intraCounter >= 2 {
			tmp := strings.Split(line, " ")
			dest, err := strconv.Atoi(tmp[0])
			if err != nil {
				panic(err)
			}
			source, err := strconv.Atoi(tmp[1])
			if err != nil {
				panic(err)
			}
			rangeL, err := strconv.Atoi(tmp[2])
			if err != nil {
				panic(err)
			}
			newMap := mapNums{
				dest,
				source,
				rangeL,
			}
			*currentMap = append(*currentMap, newMap)
		}
	}

	// find the location for each seed
	for _, seedNum := range seeds {
		fmt.Printf("Processing seed: %d\n", seedNum)
		soilNum := process(seedNum, seedToSoil)
		fertilizerNum := process(soilNum, soilToFertilizer)
		waterNum := process(fertilizerNum, fertilizerToWater)
		lightNum := process(waterNum, waterToLight)
		tempNum := process(lightNum, lightToTemp)
		humidityNum := process(tempNum, tempToHumidity)
		locationNum := process(humidityNum, humidityToLocation)

		if locationNum < res {
			res = locationNum
		}
	}

	fmt.Printf("FINAL RESULT: %d\n", res)
}

// based on the provided map m, returns the correct destination number for the provided n
func process(n int, maps []mapNums) int {
	for _, m := range maps {
		start := m.sourceStart
		end := m.sourceStart + m.rangeLen - 1
		if n >= start && n <= end {
			return n - start + m.destStart
		}
	}
	return n
}
