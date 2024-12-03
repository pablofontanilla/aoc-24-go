package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"regexp"
	"strconv"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {

	// Compile the regular expression
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Find all matches
	matches := re.FindAllStringSubmatch(input, -1)

	totalSum := 0
	// Iterate through the matches
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		totalSum += num1 * num2
	}
	if part2 {
		return "42"
	}
	// solve part 1 here
	return totalSum
}
