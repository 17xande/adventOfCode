package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type move struct {
	amount      int
	source      int
	destination int
}

type stacks [9][]rune

func main() {
	rawStacks, rawMoves := input()
	stacks := parseStacks(rawStacks)
	moves := parseMoves(rawMoves)

	for _, m := range moves {
		// fmt.Printf("\nmove no: %d (%v)\n", i, m)
		// for j, s := range stacks {
		// 	fmt.Printf("%d: %+v\n", j, string(s))
		// }
		stacks.move(m)
	}

	for j, s := range stacks {
		fmt.Printf("%d: %+v\n", j, string(s))
	}

	for _, s := range stacks {
		l := len(s) - 1
		if l == -1 {
			fmt.Printf(" ")
			continue
		}
		fmt.Printf("%s", string(s[l]))
	}
}

func (s *stacks) move(m move) {
	for i := 0; i < m.amount; i++ {
		sourceStack := s[m.source]
		lastIndex := len(sourceStack) - 1
		if lastIndex < 0 {
			fmt.Printf("something went wrong:\n%v\n", m)
			for j, s := range s {
				fmt.Printf("%d: %+v\n", j, string(s))
			}
		}
		c := sourceStack[lastIndex]
		s[m.destination] = append(s[m.destination], c)
		s[m.source] = sourceStack[:lastIndex]
	}
}

func parseStacks(rawStacks []string) stacks {
	// There are 9 stacks, 8 crates high
	stacks := stacks{}

	for _, line := range rawStacks {
		for i, j := 1, 0; i < 35; i, j = i+4, j+1 {
			crate := []rune(line)[i]
			if string(crate) == " " {
				continue
			}
			stacks[j] = append(stacks[j], crate)
		}
	}

	// Reverse the crates in each stack.
	for _, s := range stacks {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	return stacks
}

func parseMoves(rawMoves []string) []move {
	moves := []move{}

	for _, rm := range rawMoves {
		arr := strings.Split(rm, " ")
		amount, _ := strconv.Atoi(arr[1])
		source, _ := strconv.Atoi(arr[3])
		destination, _ := strconv.Atoi(arr[5])
		m := move{
			amount:      amount,
			source:      source - 1,
			destination: destination - 1,
		}
		moves = append(moves, m)
	}

	return moves
}

func input() (rawStacks []string, rawMoves []string) {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	rawStacks = lines[:8]
	rawMoves = lines[10:]
	return rawStacks, rawMoves
}
