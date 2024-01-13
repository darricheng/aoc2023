package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
1. Slice of length 2 arrays, each array is [start, range].
2. At each step (e.g. changing from soil to water), there are two sub-steps.
3. Split & transform. For each array, split and transform the numbers according to the rules for that step. This step can be done in parallel.
4. Sort & merge. Sort the new slice of arrays from smallest to biggest according to the start number. Merge any overlapping arrays.
5. Repeat steps 3-4 until the location numbers are found. Take slice[0][0] to get the smallest location.
*/

type mapNums struct {
	destStart   int
	sourceStart int
	rangeLen    int
}

var seedRanges [][2]int // slice of arrays. where each array is [start. range]
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
			seedRangesStr := strings.Split(line, " ")[1:] // discard first val which is "seeds:"
			var seedRangesNum []int
			// Convert strings to int
			for _, s := range seedRangesStr {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				seedRangesNum = append(seedRangesNum, n)
			}
			for j := 0; j < len(seedRangesNum); j += 2 {
				s := seedRangesNum[j]
				r := seedRangesNum[j+1]
				seedRanges = append(seedRanges, [2]int{s, r})
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

	soilRanges := process(seedRanges, seedToSoil)
	fertilizerRanges := process(soilRanges, soilToFertilizer)
	waterRanges := process(fertilizerRanges, fertilizerToWater)
	lightRanges := process(waterRanges, waterToLight)
	tempRanges := process(lightRanges, lightToTemp)
	humidityRanges := process(tempRanges, tempToHumidity)
	locationRanges := process(humidityRanges, humidityToLocation)

	fmt.Printf("FINAL RESULT: %d\n", res)
}

// based on the provided map m, returns the correct destination number for the provided n
func process(n [][2]int, maps []mapNums) [][2]int {
	/*
		# Split and transform

		We loop through the list of ranges, applying every map to it to transform the numbers.
		If the range gets split in half, we append one half to the end of the ranges slice
		and continue with the other half.
		Once we go through all the maps for a range, any leftover range will just be the same

		There are 6 different possibilities for each range to map pair:

		1,2. The range is completely to the left or right of the map
		range: |-----|    or    |----|
		map:            |-----|
		3,4. The range intersects the map at the left or right edge
		range: |-----|  or  |--------|
		map:      |-------------|
		5. The range is a subset of the map
		range:      |---|
		map:     |-----------|
		6. The range is a superset of the map
		range: |-------------|
		map:        |-----|

		For ranges that fall right at the edge of the map, go by the following thinking

		3,4. There is at least one value that falls outside one side of the map
		6. There is at least one value that falls outside each side of the map

		For each possibility, the steps are subtly different
		1,2. We skip the map entirely
		3,4. Shrink the range, then append the transformed range to the slice of transformed ranges
		5. Transform the entire range, continue with the next range in the list
		6. Append the transformed range, add the rear to the list of ranges to transform,
		   continue with the front

		If the range goes through all the maps and there is leftover range, add it to
		the slice of transformed ranges

		The list of ranges will be treated as a queue, as in case 6,
		we enqueue another pair to check later
	*/
	var transformedRanges [][2]int

	for {
		// Keep iterating until all range pairs are removed from the queue
		if len(n) == 0 {
			break
		}
		// Dequeue the first range pair
		rangePair := n[0]
		n = n[1:]

		// loop through the maps to compare against the pair
		for _, mapPair := range maps {
			// check which case 1-6 it is
			var caseId int
			rangeStart := rangePair[0]
			rangeEnd := rangeStart + rangePair[1] - 1 // -1 to count for the rangeStart
			mapStart := mapPair.sourceStart
			mapEnd := mapStart + mapPair.rangeLen - 1

			if rangeStart < mapStart { // 1,3,6
				if rangeEnd < mapStart {
					caseId = 1
				} else if rangeEnd <= mapEnd {
					caseId = 3
				} else {
					caseId = 6
				}
			} else { // 2,4,5
				if rangeStart > mapEnd {
					caseId = 2
				} else if rangeEnd <= mapEnd {
					caseId = 5
				} else {
					caseId = 4
				}
			}

			leftoverStart := rangeStart
			leftoverEnd := rangeEnd
			switch caseId {
			case 1: // do nothing
			case 2: // do nothing
			case 3:
				leftoverEnd = mapStart - 1
				transformedLen := rangeEnd - leftoverEnd
				transformedRange := [2]int{mapPair.destStart, transformedLen}
				transformedRanges = append(transformedRanges, transformedRange)
			case 4:
				leftoverStart = mapEnd + 1
				transformedLen := leftoverStart - rangeStart
				transformedStart := rangeStart - mapStart + mapPair.destStart
				transformedRange := [2]int{transformedStart, transformedLen}
				transformedRanges = append(transformedRanges, transformedRange)
			case 5:
				transformedStart := rangeStart - mapStart + mapPair.destStart
				transformedRange := [2]int{transformedStart, rangePair[1]}
				transformedRanges = append(transformedRanges, transformedRange)
				// break out of loop as we're done with this range
				break
			case 6:
				leftoverEnd = mapStart - 1
				transformedRange := [2]int{mapPair.destStart, mapPair.rangeLen}
				transformedRanges = append(transformedRanges, transformedRange)
				// append leftover on right side to queue
				extraStart := mapEnd + 1
				extraLen := rangeEnd - mapEnd
				extra := [2]int{extraStart, extraLen}
				n = append(n, extra)
			}
			rangePair[0] = leftoverStart
			rangePair[1] = leftoverEnd - leftoverStart + 1 // +1 to include the start value
		}
		// finished checking through all the maps for the current range pair
		// Add the leftover range to the transformed ranges slice
		transformedRanges = append(transformedRanges, rangePair)
	}
	// done with checking all the range pairs
	// TODO: I think I'm done with the split and transform implementation, need to test

	/*
		# Sort and merge

		From split and transform, we should end up with a new slice of slices that
		are `[start, range]`.
		However, there might be some ranges that overlap(?), so we sort the ranges by their start,
		then merge any overlapping ranges together.
	*/

}
