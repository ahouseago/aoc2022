package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
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

	var totals []int
	total := 0
	for _, line := range strings.Split(input, "\n") {
		if strings.TrimSpace(line) == "" && total != 0 {
			totals = append(totals, total)
			total = 0
			continue
		}
		amount, _ := strconv.Atoi(line)
		total += int(amount)
	}

	sort.Ints(totals)
	fmt.Println("Part 1:", totals[len(totals)-1])

	top3 := 0
	for _, n := range totals[len(totals)-3:] {
		top3 += n
	}
	fmt.Println("Part 2:", top3)
}
