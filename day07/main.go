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

	instructions := strings.Split(strings.TrimPrefix(strings.TrimSpace(input), "$ "), "$ ")

	root := &node{
		name:     "/",
		parent:   nil,
		children: make(map[string]*node),
		size:     0,
	}

	var current *node

	for _, inst := range instructions {
		lines := strings.Split(strings.TrimSpace(inst), "\n")
		if strings.HasPrefix(lines[0], "cd ") {
			dirName := strings.TrimPrefix(lines[0], "cd ")
			switch dirName {
			case "/":
				current = root
			case "..":
				current = current.parent
			default:
				if current == nil || current.children[dirName] == nil {
					fmt.Println("Error: cannot cd into missing dir", inst)
					continue
				}
				current = current.children[dirName]
			}
		} else {
			if current == nil {
				fmt.Println("Error: current node is nil but ls called", inst)
				continue
			}
			for _, o := range lines[1:] {
				if strings.HasPrefix(o, "dir ") {
					dirName := strings.TrimPrefix(o, "dir ")
					current.append(dirName, 0)
				} else {
					parts := strings.Split(o, " ")
					size, _ := strconv.Atoi(parts[0])
					current.append(parts[1], size)
				}
			}
		}
	}

	dirs, totalSize, totalUnder100k := root.calculateSize(100000)
	fmt.Println("Part 1:", totalUnder100k)

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].size < dirs[j].size
	})
	for _, d := range dirs {
		if d.size+70000000-totalSize > 30000000 {
			fmt.Println("Part 2:", d.size)
			break
		}
	}
}

type node struct {
	name     string
	size     int
	parent   *node
	children map[string]*node
}

func (n *node) append(name string, size int) {
	if c := n.children[name]; c != nil {
		return
	}
	n.children[name] = &node{
		name:     name,
		size:     size,
		parent:   n,
		children: make(map[string]*node),
	}
}

func (n *node) calculateSize(maxDirSize int) ([]*node, int, int) {
	if len(n.children) == 0 {
		return nil, n.size, 0
	}
	dirs, size, underTotal := []*node{}, n.size, 0
	for _, c := range n.children {
		d, cSize, under := c.calculateSize(maxDirSize)
		size += cSize
		underTotal += under
		dirs = append(dirs, d...)
	}
	if size < maxDirSize {
		underTotal += size
	}
	dirs = append(dirs, &node{name: n.name, size: size})
	return dirs, size, underTotal
}
