package main

import (
	"github.com/bitfield/script"
	"github.com/jpillora/puzzler/harness/aoc"
	"sort"
	"strconv"
)

func main() {
	aoc.Harness(run)
}

func stringToIntSlice(strings []string) []int {

	integers := make([]int, len(strings))

	for i, s := range strings {
		integers[i], _ = strconv.Atoi(s)
	}
	return integers
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func getNumTimesInList(frequencyColumn []int, orderedColumn []int, searchedNumber int) int {

	for i := range frequencyColumn {
		if orderedColumn[i] == searchedNumber {
			return frequencyColumn[i]
		}
	}
	return 0
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block

	leftNumbersString, _ := script.Echo(input).Column(1).Slice()
	rightNumbersString, _ := script.Echo(input).Column(2).Slice()
	leftNumbers := stringToIntSlice(leftNumbersString)
	rightNumbers := stringToIntSlice(rightNumbersString)

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	if part2 {

		var totalSimilarity int = 0

		var rightColumnFrequenciesString, _ = script.Echo(input).Column(2).Freq().Column(1).Slice()
		var rightColumnNumbersString, _ = script.Echo(input).Column(2).Freq().Column(2).Slice()
		var rightColumnFrequencies = stringToIntSlice(rightColumnFrequenciesString)
		var rightColumnNumbers = stringToIntSlice(rightColumnNumbersString)

		for i := range leftNumbers {
			// index is the index where we are
			// element is the element from someSlice for where we are
			var number = leftNumbers[i]
			var timesNumber = getNumTimesInList(rightColumnFrequencies, rightColumnNumbers, number)
			totalSimilarity += number * timesNumber
		}

		return totalSimilarity
	}
	// solve part 1 here

	var totalDistance = 0
	for i := range leftNumbers {
		// index is the index where we are
		// element is the element from someSlice for where we are
		totalDistance += absDiffInt(leftNumbers[i], rightNumbers[i])
	}
	return totalDistance
}
