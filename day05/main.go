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

	inputParts := strings.Split(input, "\n\n")
	state := strings.Split(inputParts[0], "\n")
	numStacks := (len(state[len(state)-1])-3)/4 + 1

	stacks1, stacks2 := make([]stack, numStacks), make([]stack, numStacks)
	for _, line := range state[:len(state)-1] {
		for i := 1; i < len(line); i += 4 {
			if line[i] != ' ' {
				stacks1[(i-1)/4] = append(stacks1[(i-1)/4], rune(line[i]))
				stacks2[(i-1)/4] = append(stacks2[(i-1)/4], rune(line[i]))
			}
		}
	}

	for _, instruction := range strings.Split(strings.TrimSpace(inputParts[1]), "\n") {
		var x, num, from, to string
		fmt.Sscan(instruction, &x, &num, &x, &from, &x, &to)
		n, _ := strconv.Atoi(num)
		f, _ := strconv.Atoi(from)
		t, _ := strconv.Atoi(to)

		// 9000
		for i := 0; i < n; i++ {
			head, tail := stacks1[f-1][0], stacks1[f-1][1:]
			stacks1[t-1] = append(stack{head}, stacks1[t-1]...)
			stacks1[f-1] = tail
		}

		// 9001
		head := make(stack, n)
		copy(head, stacks2[f-1][:n])
		stacks2[t-1] = append(head, stacks2[t-1]...)
		stacks2[f-1] = stacks2[f-1][n:]
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

type stack []rune

func (s stack) pop() (rune, stack) {
	return s[0], s[1:]
}

func (s stack) format() string {
	out := ""
	for _, r := range s {
		out += string(r)
	}
	return out
}
