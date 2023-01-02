package main

import (
	_ "embed"
	"fmt"
	"math"
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

	var instructions []instruction
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, " ")
		val, _ := strconv.Atoi(parts[1])
		instructions = append(instructions, instruction{
			d: Direction(parts[0]),
			v: val,
		})
	}

	fmt.Println("Part 1:", solve(instructions, 2))
	fmt.Println("Part 2:", solve(instructions, 10))
}

func solve(instructions []instruction, ropeLength int) int {
	r := make([]point, ropeLength)
	visited := map[point]bool{{0, 0}: true}

	for _, inst := range instructions {
		for i := 0; i < inst.v; i++ {
			r[0].move(inst.d)
			for j := 1; j < len(r); j++ {
				head, tail := r[j-1], r[j]
				if dist(head, tail) > 1.5 {
					dir := follow(head, tail)
					r[j] = add(tail, dir)
				}
			}
			visited[r[len(r)-1]] = true
		}
	}
	return len(visited)
}

type Direction string

const (
	U Direction = "U"
	D Direction = "D"
	L Direction = "L"
	R Direction = "R"
)

type instruction struct {
	d Direction
	v int
}

type point struct {
	x, y float64
}

func (p *point) move(dir Direction) {
	switch dir {
	case "U":
		p.y++
	case "D":
		p.y--
	case "L":
		p.x--
	case "R":
		p.x++
	}
}

func add(a, b point) point {
	return point{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

func dist(a, b point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func vec(a, b point) point {
	return point{a.x - b.x, a.y - b.y}
}

// A broken version of math.Ceil that works away from 0 rather than -Infinity.
func ceil(f float64) float64 {
	if f < 0 {
		return math.Floor(f)
	}
	return math.Ceil(f)
}

func follow(p, target point) point {
	v := vec(p, target)
	v.x = ceil(v.x / 2)
	v.y = ceil(v.y / 2)
	return v
}
