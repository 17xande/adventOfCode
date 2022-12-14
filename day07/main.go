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
}

func main() {
	lines := input()
	root := parse(lines)
	root.print(0)

	// fmt.Printf("\nResult: %+v\n", parse(lines))

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

	root.ls(lines, 2)

	return &root
}

func (f *file) ls(lines []string, currentLine int) {
	for i := currentLine; i < len(lines); i++ {
		l := lines[i]
		if l == "$ cd .." {
			return
		}

		var child file
		if l[:4] == "$ cd" {
			f = f.children[l[5:]]
			continue
		}

		if l[:4] == "$ ls" {
			f.ls(lines, i+1)
			break
		}

		if l[:3] == "dir" {
			child = file{
				name:     l[4:],
				isDir:    true,
				children: make(map[string]*file),
			}
		} else {
			s := strings.Split(l, " ")
			size, _ := strconv.Atoi(s[0])
			child = file{
				name:  s[1],
				isDir: false,
				size:  size,
			}
		}

		f.children[child.name] = &child
	}
	f.sumSize()
}

func input() []string {
	bytes, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bytes), "\n")
}
