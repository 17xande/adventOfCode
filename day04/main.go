package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sectionRange struct {
	start int
	end   int
}

func main() {
	lines := input()
	ranges := parseRanges(lines)

	cContained := 0
	cOverlapped := 0
	for _, r := range ranges {
		// fmt.Printf("%d-%d,%d-%d: contained: %t\n", r[0].start, r[0].end, r[1].start, r[1].end, contained(r))
		// fmt.Printf("%d-%d,%d-%d: overlapped: %t\n", r[0].start, r[0].end, r[1].start, r[1].end, overlapped(r))
		if contained(r) {
			cContained++
		}
		if overlapped(r) {
			cOverlapped++
		}
	}

	fmt.Printf("%d assignment pairs have a range that fully contains the other\n", cContained)
	fmt.Printf("%d assignment pairs have a range that overlaps the other\n", cOverlapped)
}

func overlapped(sr [2]sectionRange) bool {
	if sr[0].start < sr[1].start && sr[0].end < sr[1].start {
		return false
	}
	if sr[0].end > sr[1].end && sr[0].start > sr[1].end {
		return false
	}

	return true
}

func contained(sr [2]sectionRange) bool {
	if sr[0].start <= sr[1].start && sr[0].end >= sr[1].end {
		return true
	}
	if sr[1].start <= sr[0].start && sr[1].end >= sr[0].end {
		return true
	}
	return false
}

func parseRanges(lines []string) [][2]sectionRange {
	result := [][2]sectionRange{}
	for _, l := range lines {
		s := strings.Split(l, ",")
		r11, _ := strconv.Atoi(strings.Split(s[0], "-")[0])
		r12, _ := strconv.Atoi(strings.Split(s[0], "-")[1])
		r21, _ := strconv.Atoi(strings.Split(s[1], "-")[0])
		r22, _ := strconv.Atoi(strings.Split(s[1], "-")[1])
		range1 := sectionRange{
			start: r11,
			end:   r12,
		}
		range2 := sectionRange{
			start: r21,
			end:   r22,
		}

		r := [2]sectionRange{range1, range2}

		result = append(result, r)
	}

	return result
}

func input() []string {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	return lines
}
