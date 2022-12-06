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

	inputParts := strings.Split(input, "\n\n")
	state := strings.Split(inputParts[0], "\n")
	numStacks := (len(state[len(state)-1])-3)/4 + 1

	stacks1, stacks2 := make([][]rune, numStacks), make([][]rune, numStacks)
	for _, line := range state[:len(state)-1] {
		for i := 1; i < len(line); i += 4 {
			if line[i] != ' ' {
				stacks1[(i-1)/4] = append(stacks1[(i-1)/4], rune(line[i]))
				stacks2[(i-1)/4] = append(stacks2[(i-1)/4], rune(line[i]))
			}
		}
	}

	for _, instruction := range strings.Split(strings.TrimSpace(inputParts[1]), "\n") {
		var n, from, to int
		fmt.Sscanf(instruction, "move %d from %d to %d\n", &n, &from, &to)
		from, to = from-1, to-1

		// 9000
		for i := 0; i < n; i++ {
			head, tail := stacks1[from][0], stacks1[from][1:]
			stacks1[to] = append([]rune{head}, stacks1[to]...)
			stacks1[from] = tail
		}

		// 9001
		head := make([]rune, n)
		copy(head, stacks2[from][:n])
		stacks2[to] = append(head, stacks2[to]...)
		stacks2[from] = stacks2[from][n:]
	}

	part1 := ""
	for _, s := range stacks1 {
		part1 += string(s[0])
	}
	part2 := ""
	for _, s := range stacks2 {
		part2 += string(s[0])
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
