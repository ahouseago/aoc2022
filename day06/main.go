package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed input.txt
var input string

var (
	test1 = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	test2 = "bvwbjplbgvbhsrlpgdmjqwftvncz"
	test3 = "nppdvjthqldpwncqszvftbrmjlhg"
	test4 = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	test5 = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
)

func main() {
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "test") {
		switch os.Args[1] {
		case "test", "test1":
			input = test1
		case "test2":
			input = test2
		case "test3":
			input = test3
		case "test4":
			input = test4
		case "test5":
			input = test5
		}
	}

	fmt.Println("Part 1:", findSequence(4))
	fmt.Println("Part 2:", findSequence(14))
}

func findSequence(length int) int {
	for i := range []rune(strings.TrimSpace(input)) {
		if i < length {
			continue
		}
		duplicate, matches := false, make(map[rune]bool, length)
		for _, r := range []rune(input)[i-length : i] {
			if matches[r] == true {
				duplicate = true
				break
			}
			matches[r] = true
		}
		if !duplicate {
			return i
		}
	}

	return -1
}
