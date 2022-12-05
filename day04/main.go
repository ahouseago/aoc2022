package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test_input.txt
var testInput string

func main() {
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = testInput
	}

	part1 := 0
	part2 := 0
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		elves := strings.Split(line, ",")
		larger, smaller := toRange(elves[0]), toRange(elves[1])
		if len(smaller) > len(larger) {
			larger, smaller = smaller, larger
		}
		if contains(larger, smaller) {
			part1++
		}
		if overlaps(larger, smaller) {
			part2++
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func toRange(s string) []int {
	vals := strings.Split(s, "-")
	start, _ := strconv.Atoi(vals[0])
	end, _ := strconv.Atoi(vals[1])
	out := make([]int, end-start+1)
	for i := start; i <= end; i++ {
		out[i-start] = i
	}
	return out
}

func contains(larger, smaller []int) bool {
	return larger[0] <= smaller[0] && larger[len(larger)-1] >= smaller[len(smaller)-1]
}

func overlaps(larger, smaller []int) bool {
	l0, l1, s0, s1 := larger[0], larger[len(larger)-1], smaller[0], smaller[len(smaller)-1]
	return contains(larger, smaller) || (l0 <= s0 && l1 >= s0) || (l0 <= s1 && l1 >= s1)
}
