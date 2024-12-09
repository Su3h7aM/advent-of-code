package main

import "core:fmt"
import "core:os"
import "core:slice"
import "core:strconv"
import "core:strings"

main :: proc() {
	input := input_read("2024/inputs/01.txt")
	defer delete(input)

	left, right := parse_input(input)
	defer delete(left)
	defer delete(right)

	p1 := part_1(left, right)
	p2 := part_2(left, right)

	fmt.println("AOC day 01")
	fmt.printfln("Part One: %v", p1)
	fmt.printfln("Part Two: %v", p2)
}

part_1 :: proc(left, right: []int) -> (total_diff: int) {
	slice.sort(left[:])
	slice.sort(right[:])

	for _, i in left {
		total_diff += abs(left[i] - right[i])
	}

	return
}

part_2 :: proc(left, right: []int) -> (score: int) {
	rept := make(map[int]int, len(left))
	defer delete(rept)

	for v in right {
		rept[v] += 1
	}

	for v in left {
		score += v * rept[v]
	}

	return
}

input_read :: proc(filename: string) -> (input: string) {
	data, ok := os.read_entire_file_from_filename(filename)
	assert(ok, "Unable to read file")

	input = string(data)

	return
}

parse_input :: proc(content: string) -> (left, right: []int) {
	lines := strings.split_lines(content)
	defer delete(lines)

	left = make([]int, len(lines))
	right = make([]int, len(lines))

	for line, i in lines {
		values := strings.split(line, "   ")
		defer delete(values)

		if len(values) != 2 {
			continue
		}

		left[i] = strconv.atoi(values[0])
		right[i] = strconv.atoi(values[1])
	}

	return
}
