package main

import (
	"testing"
)

var (
	left  = []int{3, 4, 2, 1, 3, 3}
	right = []int{4, 3, 5, 3, 9, 3}
)

func TestPart1(t *testing.T) {
	total := Part1(left, right)

	if total != 11 {
		t.Errorf("Part1 must be 31, got %d", total)
	}
}

func TestPart2(t *testing.T) {
	score := Part2(left, right)

	if score != 31 {
		t.Errorf("Part2 must be 31, got %d", score)
	}
}
