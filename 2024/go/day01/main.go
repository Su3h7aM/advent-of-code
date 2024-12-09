package main

import (
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
	var score int
	repeated := make(map[int]int)

	for _, v := range right {
		repeated[v] += 1
	}

	for _, v := range left {
		score += v * repeated[v]
	}

	return score
}

func ReadInput() (left, right []int) {
	input, err := os.ReadFile("2024/inputs/01.txt")
	if err != nil {
		fmt.Printf("Failed to open file: %s", err)
	}

	lines := strings.Split(string(input), "\n")

	leftColumn := make([]int, len(lines))
	rightColumn := make([]int, len(lines))

	for i, line := range lines {
		columns := strings.Fields(line)

		if len(columns) != 2 {
			continue
		}

		leftColumn[i], err = strconv.Atoi(columns[0])
		if err != nil {
			fmt.Printf("string to int error: %s", err)
		}

		rightColumn[i], err = strconv.Atoi(columns[1])
		if err != nil {
			fmt.Printf("string to int error: %s", err)
		}
	}

	return leftColumn, rightColumn
}
