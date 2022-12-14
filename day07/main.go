package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type file struct {
	name     string
	size     int
	isDir    bool
	children map[string]*file
	parent   *file
}

func main() {
	lines := input()
	root := parse(lines)
	root.sumSize()
	// root.print(0)
	sum := sumDirUnder(root, 100000)
	fmt.Printf("Sum: %d\n", sum)

	available := 70000000 - root.size
	required := 30000000 - available
	fmt.Printf("\nCurrent available space: %d, require additional space of: %d\n", available, required)

	under := findDirUnder(root, 30000000, required)
	fmt.Printf("Found dir %d large to delete\n", under)
}

func findDirUnder(f *file, min, limit int) int {
	for _, c := range f.children {
		if !c.isDir {
			continue
		}

		if c.size >= limit && c.size < min {
			min = c.size
		}

		min = findDirUnder(c, min, limit)
	}

	return min
}

func sumDirUnder(f *file, limit int) int {
	total := 0

	for _, c := range f.children {
		if !c.isDir {
			continue
		}

		if c.size <= limit {
			total += c.size
		}

		total += sumDirUnder(c, limit)
	}

	return total
}

func (f *file) print(indent int) {
	if f.isDir {
		fmt.Printf("%s-d %s %d\n", strings.Repeat(" ", indent), f.name, f.size)

		for _, c := range f.children {
			c.print(indent + 2)
		}
	} else {
		fmt.Printf("%s-f %s %d\n", strings.Repeat(" ", indent), f.name, f.size)
	}
}

func (f *file) sumSize() {
	if !f.isDir {
		return
	}

	for _, c := range f.children {
		f.size += c.size
	}
}

func parse(lines []string) *file {
	root := file{
		name:     "/",
		isDir:    true,
		children: make(map[string]*file),
	}

	currentDir := &root

	for _, l := range lines[1:] {
		if l == "$ cd .." {
			currentDir.sumSize()
			currentDir = currentDir.parent
			continue
		}

		if l[:4] == "$ cd" {
			currentDir = currentDir.children[l[5:]]
			continue
		}

		if l[:4] == "$ ls" {
			continue
		}

		var child file
		if l[:3] == "dir" {
			child = file{
				name:     l[4:],
				isDir:    true,
				children: make(map[string]*file),
				parent:   currentDir,
			}
		} else {
			s := strings.Split(l, " ")
			size, _ := strconv.Atoi(s[0])
			child = file{
				name:   s[1],
				isDir:  false,
				size:   size,
				parent: currentDir,
			}
		}

		currentDir.children[child.name] = &child
	}
	currentDir.sumSize()
	return &root
}

func input() []string {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bytes), "\n")
}
