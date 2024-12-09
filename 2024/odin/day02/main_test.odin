package main

import "core:testing"

result_01 :: 2
result_02 :: 4

@(test)
test_part01 :: proc(t: ^testing.T) {
	input := input_read("2024/inputs/02_sample.txt")

	reports := parse_input(&input)
	defer delete(reports)

	p1 := part01(reports)

	testing.expectf(t, p1 == result_01, "Part One: %v", p1)
}

@(test)
test_part02 :: proc(t: ^testing.T) {
	input := input_read("2024/inputs/02_sample.txt")

	reports := parse_input(&input)
	defer delete(reports)

	new_reports := part02(reports)
	p2 := part01(new_reports)

	testing.expectf(t, p2 == result_02, "Part Two: %v", p2)
}
