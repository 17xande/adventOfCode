package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type forest [][]int

func main() {
	lines := input()
	forest := parse(lines)
	// x := 1
	// y := 3
	// v := forest.treeVisible(x, y)

	// fmt.Printf("\nTree (%d;%d) %d visible: %t\n", x, y, forest[y][x], v)

	count := forest.count()

	fmt.Printf("%d trees are visible\n", count)
}

func (f *forest) count() int {
	count := 0
	fo := *f

	for y := 1; y < len(fo)-1; y++ {
		for x := 1; x < len(fo[0])-1; x++ {
			if f.treeVisible(x, y) {
				count++
			}
			// fmt.Printf("Tree (%d;%d) %d visible: %t\n", x, y, fo[y][x], f.treeVisible(x, y))
		}
	}

	count += (len(fo) - 1) * 2
	count += (len(fo[0]) - 1) * 2

	return count
}

func (f *forest) treeVisible(x, y int) bool {
	fore := *f
	tree := fore[y][x]

	visLeft := true
	for _, t := range fore[y][:x] {
		// fmt.Printf("%d", t)
		if t >= tree {
			visLeft = false
			break
		}
	}
	if visLeft {
		return true
	}

	visRight := true
	for _, t := range fore[y][x+1:] {
		// fmt.Printf("%d", t)
		if t >= tree {
			visRight = false
			break
		}
	}
	if visRight {
		return true
	}

	visTop := true
	for _, row := range fore[:y] {
		t := row[x]
		// fmt.Printf("%d\n", t)
		if t >= tree {
			visTop = false
			break
		}
	}
	if visTop {
		return true
	}

	visBot := true
	for _, row := range fore[y+1:] {
		t := row[x]
		if t >= tree {
			visBot = false
		}
	}

	// if x == 1 && y == 3 {
	// 	fmt.Printf("(1;3) - %d t:%t,b:%t,l:%t,r:%t\n", tree, visTop, visBot, visLeft, visRight)
	// 	fmt.Printf("%v\n", fore[y+1:])
	// }
	return visBot

}

func parse(lines []string) forest {
	width := len(lines[0])
	breadth := len(lines)
	forest := make([][]int, breadth)
	for i := range forest {
		forest[i] = make([]int, width)
	}

	for i, l := range lines {
		for j, r := range l {
			n, _ := strconv.Atoi(string(r))
			forest[i][j] = n
		}
	}

	return forest
}

func input() []string {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bytes), "\n")
}
