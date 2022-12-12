package main

import (
	"fmt"
	"os"
	"strings"
)

type move struct {
	amount      int
	source      int
	destination int
}

type stack [8]rune

func main() {
	rawStacks, rawMoves := input()

	fmt.Printf("%s\n%s", rawStacks, rawMoves)
}

func parseStacks(rawStacks []string) [9]stack {
	stacks := [9]stack{}

	return stacks
}

func parseMoves(rawMoves []string) []move {
	moves := []move{}

	return moves
}

func input() (rawStacks []string, rawMoves []string) {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	rawStacks = lines[:8]
	rawMoves = lines[11:]
	return rawStacks, rawMoves
}
