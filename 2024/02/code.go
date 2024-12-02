package main

import (
	_ "github.com/bitfield/script"
	"github.com/jpillora/puzzler/harness/aoc"
	"strconv"
	"strings"
)

func main() {
	aoc.Harness(run)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func stringToIntSlice(strings []string) []int {

	integers := make([]int, len(strings))
	for i, line := range strings {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil
		}
		integers[i] = n
	}
	return integers
}

func isIncreasing(levels []int) bool {
	var previousValue = levels[0]
	for i := 1; i < len(levels); i++ {
		if levels[i] > previousValue {
			previousValue = levels[i]
			continue
		} else {
			return false
		}
	}
	return true
}
func isDecreasing(levels []int) bool {
	var previousValue = levels[0]
	for i := 1; i < len(levels); i++ {
		if levels[i] < previousValue {
			previousValue = levels[i]
			continue
		} else {
			return false
		}
	}
	return true
}
func differsBy(min int, max int, levels []int) bool {
	var previousValue = levels[0]
	for i := 1; i < len(levels); i++ {
		var diff = absDiffInt(previousValue, levels[i])
		if diff >= 1 && diff <= 3 {
			previousValue = levels[i]
			continue
		} else {
			return false
		}
	}
	return true
}

func isSafe(levels []string) bool {
	var levelsInt = stringToIntSlice(levels)
	var result = (isIncreasing(levelsInt) || isDecreasing(levelsInt)) && differsBy(1, 3, levelsInt)
	return result
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	var safeTotal = 0
	for _, line := range strings.Split(input, "\n") {
		splitLine := strings.Fields(line)
		if isSafe(splitLine) {
			safeTotal += 1
		}
	}
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return safeTotal
}
