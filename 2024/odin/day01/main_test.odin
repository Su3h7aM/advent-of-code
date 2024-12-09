package main

import "core:testing"

result_01 :: 11
result_02 :: 31


@(test)
test_part_01 :: proc(t: ^testing.T) {
	input := input_read("2024/inputs/01_sample.txt")
	defer delete(input)
	left, right := parse_input(input)
	defer delete(left)
	defer delete(right)

	p1 := part_1(left, right)

	testing.expectf(t, p1 == result_01, "Part One: %v", p1)
}

@(test)
test_part_02 :: proc(t: ^testing.T) {
	input := input_read("2024/inputs/01_sample.txt")
	defer delete(input)
	left, right := parse_input(input)
	defer delete(left)
	defer delete(right)

	p2 := part_2(left, right)

	testing.expectf(t, p2 == result_02, "Part Two: %v", p2)
}
