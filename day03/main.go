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

	fmt.Printf("Sack priority sum: %d\n", sum)

	groups := parseGroups(&sacks)

	sum = 0
	for _, g := range groups {
		findBadge(&g)
		// fmt.Printf("Group %d: %s %s %s Badge: %c Priority: %d\n", i, g.sacks[0].whole, g.sacks[1].whole, g.sacks[2].whole, g.badge, g.priority)
		sum += g.priority
	}

	fmt.Printf("Group priority sum: %d\n", sum)
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

func parseGroups(sacks *[]sack) []group {
	groups := []group{}
	g := group{
		sacks: [3]sack{},
	}

	for i, s := range *sacks {
		if i != 0 && i%3 == 0 {
			groups = append(groups, g)
			g = group{
				sacks: [3]sack{},
			}
		}

		g.sacks[i%3] = s
	}

	groups = append(groups, g)
	return groups
}

func findBadge(g *group) {
	for _, r := range g.sacks[0].whole {
		if strings.Contains(g.sacks[1].whole, string(r)) && strings.Contains(g.sacks[2].whole, string(r)) {
			g.badge = r
			g.priority = priorities[r]
			break
		}
	}
}

func input() []string {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	return lines
}
