package main

import "core:fmt"
// import "core:mem"
import "core:os"
import "core:strconv"
import "core:strings"

report :: struct {
	levels: []int,
	safe:   bool,
}

main :: proc() {

	// track: mem.Tracking_Allocator
	// mem.tracking_allocator_init(&track, context.allocator)
	// context.allocator = mem.tracking_allocator(&track)
	//
	// defer {
	// 	if len(track.allocation_map) > 0 {
	// 		fmt.eprintf("=== %v allocations not freed: ===\n", len(track.allocation_map))
	// 		for _, entry in track.allocation_map {
	// 			fmt.eprintf("- %v bytes @ %v\n", entry.size, entry.location)
	// 		}
	// 	}
	// 	if len(track.bad_free_array) > 0 {
	// 		fmt.eprintf("=== %v incorrect frees: ===\n", len(track.bad_free_array))
	// 		for entry in track.bad_free_array {
	// 			fmt.eprintf("- %p @ %v\n", entry.memory, entry.location)
	// 		}
	// 	}
	// 	mem.tracking_allocator_destroy(&track)
	// }


	input := input_read("2024/inputs/02.txt")

	reports := parse_input(&input)

	new_reports := part02(reports)

	p1 := part01(reports)
	p2 := part01(new_reports)

	fmt.println("AOC day 01")
	fmt.printfln("Part One: %v", p1)
	fmt.printfln("Part Two: %v", p2)
}

part01 :: proc(reports: []report) -> (safe_reports: int) {
	for report in reports {
		if report.safe {
			safe_reports += 1
		}
	}

	return
}

part02 :: proc(reports: []report) -> (new_reports: []report) {
	new_reports = make([]report, len(reports))
	// defer delete(new_reports)

	for rpt, i in reports {
		if rpt.safe {
			new_reports[i] = rpt
			continue
		}

		new_safe, new_levels := is_safe_with_dampener(rpt.levels)

		if new_safe {
			new_reports[i] = report{new_levels[:], new_safe}
			continue
		}

		new_reports[i] = rpt
	}

	return
}

input_read :: proc(filename: string) -> (input: string) {
	data, ok := os.read_entire_file_from_filename(filename)
	assert(ok, "Unable to read file")

	input = string(data)

	return
}

parse_input :: proc(content: ^string) -> (reports: []report) {
	// lines := strings.split_lines(content^)
	lines: [dynamic]string

	for line in strings.split_lines_iterator(content) {
		if len(line) == 0 {
			continue
		}

		append(&lines, line)
	}

	reports = make([]report, len(lines))

	for &line, i in lines {
		levels: [dynamic]int

		for field in strings.split_iterator(&line, " ") {
			append(&levels, strconv.atoi(field))
		}


		safe := is_safe(levels[:])

		reports[i] = report{levels[:], safe}

	}

	return
}

is_safe :: proc(levels: []int) -> bool {
	prev_diff: int

	for i := 1; i < len(levels); i += 1 {
		diff := levels[i] - levels[i - 1]

		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}

		if i > 1 {
			if (prev_diff > 0 && diff < 0) || (prev_diff < 0 && diff > 0) {
				return false
			}
		}

		prev_diff = diff
	}

	return true

}

is_safe_with_dampener :: proc(levels: []int) -> (safe: bool, new_levels: [dynamic]int) {
	new_levels = make([dynamic]int, len(levels) - 1)
	// defer delete(new_levels)

	for i := 0; i < len(levels); i += 1 {
		copy(new_levels[:i], levels[:i])
		copy(new_levels[i:], levels[i + 1:])

		new_safe := is_safe(new_levels[:])

		if new_safe {
			return true, new_levels
		}
	}

	return
}
