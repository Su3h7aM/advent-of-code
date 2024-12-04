package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left, right := ReadInput()

	fmt.Println("AOC day 01")

	// Part One
	fmt.Println("Part One:", Part1(left, right))

	// Part Two
	fmt.Println("Part Two:", Part2(left, right))
}

// Part One
func Part1(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	var diff []int

	for i := range left {
		rawDiff := left[i] - right[i]
		if rawDiff < 0 {
			diff = append(diff, -rawDiff)
		} else {
			diff = append(diff, rawDiff)
		}
	}

	var total int
	for _, val := range diff {
		total = total + val
	}

	return total
}

func Part2(left, right []int) int {
	var score []int
	var repeated int
	for _, lValue := range left {
		for _, rValue := range right {
			if lValue == rValue {
				repeated = repeated + 1
			}
		}

		score = append(score, repeated*lValue)
		repeated = 0
	}

	var totalScore int
	for _, val := range score {
		totalScore = totalScore + val
	}

	return totalScore
}

func ReadInput() (left, right []int) {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	var leftColumn []int
	var rightColumn []int

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)

		col1, err := strconv.Atoi(columns[0])
		col2, err := strconv.Atoi(columns[1])
		if err != nil {
			panic(err)
		}

		leftColumn = append(leftColumn, col1)
		rightColumn = append(rightColumn, col2)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return leftColumn, rightColumn
}
