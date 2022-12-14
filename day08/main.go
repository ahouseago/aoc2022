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

	width := strings.Index(input, "\n")
	fmt.Println("Part 1:", len(countTrees(strings.ReplaceAll(input, "\n", ""), width)))
	fmt.Println("Part 2:", partTwo(strings.ReplaceAll(input, "\n", ""), width))
}

func countTrees(input string, w int) map[int]bool {
	visible := make(map[int]bool, len(input))
	for i, arr := range []string{
		input,
		transpose(input, w),
		reverse(input),
		transpose(reverse(input), w),
	} {
		for r := 0; r < len(arr); r += w {
			var maxHeight byte
			for c := 0; c < w; c++ {
				index := getIndex(r+c, (i+1)%2, w)
				if i > 1 {
					index = (w * w) - index - 1
				}
				if arr[r+c] > maxHeight {
					maxHeight = arr[r+c]
					visible[index] = true
				}
			}
		}
	}
	return visible
}

func getIndex(i, j, w int) int {
	return i*j + w*(i*(1-j)%w) + i*(1-j)/w
}

func transpose(str string, w int) string {
	out := ""
	for i := 0; i < len(str); i++ {
		out += string(str[getIndex(i, 0, w)])
	}
	return out
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func partTwo(input string, w int) int {
	var maxDistance int
	for i := 0; i < len(input); i++ {
		if i < w || i > len(input)-w || i%w == 0 || (i+1)%w == 0 {
			// Skip the outside ones
			continue
		}
		seeUp, seeDown, seeLeft, seeRight := 0, 0, 0, 0
		// Search up
		for j := i - w; j > 0; j -= w {
			seeUp++
			if input[i] <= input[j] {
				break
			}
		}
		// Search down
		for j := i + w; j < len(input); j += w {
			seeDown++
			if input[i] <= input[j] {
				break
			}
		}
		// Search left
		for j := i - 1; (j+1)%w != 0; j-- {
			seeLeft++
			if input[i] <= input[j] {
				break
			}
		}
		// Search right
		for j := i + 1; j%w != 0; j++ {
			seeRight++
			if input[i] <= input[j] {
				break
			}
		}
		canSee := seeUp * seeDown * seeLeft * seeRight
		if canSee > maxDistance {
			maxDistance = canSee
		}
	}
	return maxDistance
}
