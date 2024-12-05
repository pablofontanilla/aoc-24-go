package main

import (
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
)

func main() {
	aoc.Harness(run)
}

var searchedString = []string{"X", "M", "A", "S"}

type Position struct {
	X int64
	Y int64
}

var directions = []Position{{-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}}

func getStartingPositions(character string, grid [][]string) []Position {
	var results = []Position{}
	maxX := (int64)(len(grid))
	maxY := (int64)(len(grid[0]))

	fmt.Println("Max dimensions", maxX, maxY)
	for x := range maxX {
		for y := range maxY {
			//fmt.Println("Evaluating ", x, y)
			if grid[x][y] == character {
				results = append(results, Position{x, y})
			}
		}
	}

	return results
}

func isInBounds(position Position, grid [][]string) bool {
	maxX := (int64)(len(grid))
	maxY := (int64)(len(grid[0]))
	if position.X >= 0 && position.X < maxX && position.Y >= 0 && position.Y < maxY {
		return true
	} else {
		return false
	}

}

// direction is a delta Position in which each component is -1, 0 or 1
func restOfWordInDirection(position Position, direction Position, word []string, grid [][]string) bool {
	subWord := word[1:]
	for _, letter := range subWord {
		position.X = position.X + direction.X
		position.Y = position.Y + direction.Y

		if isInBounds(position, grid) {
			if letter == grid[position.X][position.Y] {

				continue
			} else {
				return false
			}
		} else {
			//Out of bounds
			return false
		}

	}
	//If the string has
	return true
}

func findWordInCross(position Position, grid [][]string) bool {

	var diagUpLeft, diagUpRight, diagDownLeft, diagDownRight Position

	diagUpLeft.X = position.X - 1
	diagUpLeft.Y = position.Y + 1
	diagUpRight.X = position.X + 1
	diagUpRight.Y = position.Y + 1
	diagDownLeft.X = position.X - 1
	diagDownLeft.Y = position.Y - 1
	diagDownRight.X = position.X + 1
	diagDownRight.Y = position.Y - 1

	fmt.Printf("Looking for X-MAS in %d, %d", position.X, position.Y)
	fmt.Println()
	//if in bounds, we check for X-MAS
	if isInBounds(diagUpLeft, grid) && isInBounds(diagUpRight, grid) && isInBounds(diagDownLeft, grid) && isInBounds(diagDownRight, grid) {
		if ((grid[diagUpLeft.X][diagUpLeft.Y] == "M" && grid[diagDownRight.X][diagDownRight.Y] == "S") ||
			(grid[diagUpLeft.X][diagUpLeft.Y] == "S" && grid[diagDownRight.X][diagDownRight.Y] == "M")) &&
			((grid[diagUpRight.X][diagUpRight.Y] == "M" && grid[diagDownLeft.X][diagDownLeft.Y] == "S") ||
				(grid[diagUpRight.X][diagUpRight.Y] == "S" && grid[diagDownLeft.X][diagDownLeft.Y] == "M")) {
			return true

		} else {
			//NO X-MAS
			return false
		}
	} else {
		//"Out of bounds"

		fmt.Println()
		return false
	}

}

func parseLine(l string) []string {
	var line []string

	letters := strings.Split(l, "")
	for _, letter := range letters {
		line = append(line, letter)
	}
	return line
}

func parseGrid(g string) [][]string {
	var grid [][]string

	lines := strings.Split(g, "\n")

	for _, line := range lines {
		grid = append(grid, parseLine(line))
	}
	return grid
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	input = strings.TrimSpace(input)
	// Parse grid into 2D matrix
	grid := parseGrid(input)
	var totalSum = 0
	if part2 {
		APositions := getStartingPositions("A", grid)
		for _, position := range APositions {
			if findWordInCross(position, grid) {
				totalSum++
			}
		}
		return totalSum
	}

	// Get all starting positions
	startingPositions := getStartingPositions(searchedString[0], grid)

	for _, position := range startingPositions {
		for _, direction := range directions {
			if restOfWordInDirection(position, direction, searchedString, grid) {
				totalSum++
			}
		}

	}
	return totalSum
}
