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

	total, total2 := 0, 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		round := strings.Split(strings.TrimSpace(line), " ")
		if len(round) != 2 {
			panic(fmt.Sprintf("Not enough inputs in round: %v", round))
		}

		// Part 1 scoring
		s := resolve(choice(round[1]), choice(round[0]))
		total += s + int(choice(round[1]))

		// Part 2 scoring
		c := choiceFromOutcome(choice(round[0]), round[1])
		total2 += score(round[1]) + int(c)
	}

	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", total2)
}

const (
	Lose = "X"
	Draw = "Y"
	Win  = "Z"
)

func choiceFromOutcome(choice Choice, outcome string) Choice {
	switch choice {
	case Rock:
		switch outcome {
		case Lose:
			return Scissors
		case Draw:
			return Rock
		case Win:
			return Paper
		}
	case Paper:
		switch outcome {
		case Lose:
			return Rock
		case Draw:
			return Paper
		case Win:
			return Scissors
		}
	case Scissors:
		switch outcome {
		case Lose:
			return Paper
		case Draw:
			return Scissors
		case Win:
			return Rock
		}
	}
	return Rock
}

// How much does choice a score given choice b?
func resolve(a, b Choice) int {
	switch a {
	case Rock:
		switch b {
		case Rock:
			return 3
		case Paper:
			return 0
		case Scissors:
			return 6
		}
	case Paper:
		switch b {
		case Rock:
			return 6
		case Paper:
			return 3
		case Scissors:
			return 0
		}
	case Scissors:
		switch b {
		case Rock:
			return 0
		case Paper:
			return 6
		case Scissors:
			return 3
		}
	}
	return 0
}

type Choice int

const (
	Rock Choice = iota + 1
	Paper
	Scissors
)

func choice(letter string) Choice {
	switch letter {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	}
	fmt.Println("Got an unknown choice:", letter)
	return Rock
}

func score(letter string) int {
	switch letter {
	case Lose:
		return 0
	case Draw:
		return 3
	case Win:
		return 6
	}
	fmt.Println("Got an unknown outcome:", letter)
	return 0
}
