package main

import (
	"fmt"
	"os"
	"strings"
)

type file struct {
	name     string
	size     int
	isDir    bool
	children []file
}

func main() {
	lines := input()

	fmt.Println(lines)
}

func input() []string {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bytes), "\n")
}
