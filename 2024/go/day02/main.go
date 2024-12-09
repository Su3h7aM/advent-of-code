package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type report struct {
	levels []int
	safe   bool
}

func main() {
	input := readInput("2024/inputs/02.txt")

	reports := parseInput(input)

	safeReports := part01(reports)
	newReports := part02(reports)
	newSafeReports := part01(newReports)

	fmt.Println("AOC day 02")
	fmt.Println("Part One:", safeReports)
	fmt.Println("Part Two:", newSafeReports)
}

func part01(reports []report) int {
	var safeReports int
	for _, report := range reports {
		if report.safe {
			safeReports += 1
		}
	}

	return safeReports
}

func part02(reports []report) []report {
	var correctedReports []report

	for _, rpt := range reports {
		if rpt.safe {
			correctedReports = append(correctedReports, rpt)
			continue
		}

		newSafe, newLevels := isSafeWithDampener(rpt.levels)

		if newSafe {
			correctedReports = append(correctedReports, report{
				levels: newLevels,
				safe:   newSafe,
			})
			continue
		}

		correctedReports = append(correctedReports, rpt)
	}

	return correctedReports
}

func parseInput(input string) []report {
	lines := strings.Split(string(input), "\n")

	var reports []report

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) == 0 {
			break
		}

		levels := make([]int, len(fields))

		for i, field := range fields {
			var err error
			levels[i], err = strconv.Atoi(field)
			if err != nil {
				fmt.Printf("string to int error: %s", err)
			}
		}

		safe := isSafe(levels)
		reports = append(reports, report{levels, safe})
	}

	return reports
}

func isSafe(levels []int) bool {
	var prevDiff int

	for i := 1; len(levels) > i; i++ {
		diff := levels[i] - levels[i-1]

		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}

		if i > 1 {
			if (prevDiff > 0 && diff < 0) || (prevDiff < 0 && diff > 0) {
				return false
			}
		}

		prevDiff = diff
	}

	return true
}

func isSafeWithDampener(levels []int) (bool, []int) {
	newLevels := make([]int, len(levels)-1)

	for i := 0; i < len(levels); i++ {
		copy(newLevels[:i], levels[:i])
		copy(newLevels[i:], levels[i+1:])
		newSafe := isSafe(newLevels)

		if newSafe {
			return true, newLevels
		}
	}

	return false, nil
}

func readInput(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("failed to open file: %s", err)
	}

	return string(input)
}
