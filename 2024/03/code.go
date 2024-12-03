package main

import (
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
	"regexp"
	"strconv"
)

func main() {
	aoc.Harness(run)
}

func getSubResult(input string) int {
	subTotalSum := 0
	// Compile the regular expression
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Find all matches
	matches := re.FindAllStringSubmatch(input, -1)
	// Iterate through the matches
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		subTotalSum += num1 * num2
	}
	return subTotalSum
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {

	totalSum := getSubResult(input)
	var part2TotalSum = 0
	if part2 {
		//We get all the validated tuples (start with do or start-of-line and finish with don't())
		re2 := regexp.MustCompile(`((do\(\)|^)(.+?)(don't\(\)))`)

		negativeMatches := re2.FindAllStringSubmatch(input, -1)
		fmt.Printf("Found %v matches\n", len(negativeMatches))

		//Calculate and accumulate the subtotal from each match's tuples
		for _, match := range negativeMatches {
			part2TotalSum += getSubResult(match[3])
		}

		return part2TotalSum
	}
	// solve part 1 here
	return totalSum
}
