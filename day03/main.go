package main

import (
	_ "embed"
	"fmt"
	"os"
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

	total := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		first, second := set(line[:len(line)/2]), set(line[len(line)/2:])
		for l := range first {
			if second[l] {
				total += int(value([]rune(l)[0]))
				break
			}
		}
	}

	total2 := 0
	for i := 0; i < len(lines); i += 3 {
		items := make(map[string]int)
		for j := i; j < i+3; j++ {
			for l := range set(lines[j]) {
				items[l]++
			}
		}
		for l, item := range items {
			if item == 3 {
				total2 += int(value([]rune(l)[0]))
			}
		}
	}

	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", total2)
}

func set(input string) map[string]bool {
	out := make(map[string]bool, len(input))
	for _, l := range strings.Split(input, "") {
		out[l] = true
	}
	return out
}

func value(r rune) rune {
	if strings.ToLower(string(r)) == string(r) {
		return r - 96
	}
	return r - 38
}
