package main

import (
	"fmt"
	"os"
	"strings"
)

type sack struct {
	whole        string
	compartment1 string
	compartment2 string
	duplicate    string
	priority     int
}

type group = struct {
	sacks    [3]sack
	badge    rune
	priority int
}

var priorities = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}

func main() {
	lines := input()
	sacks := parse(lines)
	// fmt.Printf("%+v\n", sacks[:10])

	sum := 0
	for _, s := range sacks {
		sum += s.priority
	}

	fmt.Printf("Priority sum: %d\n", sum)
}

func parse(lines []string) []sack {
	sacks := []sack{}

	for i, line := range lines {
		h := len(line) / 2
		s := sack{
			whole:        line,
			compartment1: line[:h],
			compartment2: line[h:],
		}

		if len(s.compartment1) != len(s.compartment2) {
			panic("different length compartments for line: " + fmt.Sprint(i))
		}

		if len(s.compartment1)+len(s.compartment2) != len(line) {
			panic("compartment lengths don't add up to line length for line: " + fmt.Sprint(i))
		}

		findPriority(&s)

		sacks = append(sacks, s)
	}

	return sacks
}

func findPriority(sack *sack) {
	for _, r := range []rune(sack.compartment1) {
		if strings.Contains(sack.compartment2, string(r)) {
			sack.duplicate = string(r)
			sack.priority = priorities[r]
			break
		}
	}
}

func findBadge(sacks *[]sack) {

}

func input() []string {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	return lines
}
