package main

import (
	"os"
	"strconv"
	"strings"
)

type move struct {
	direction string
	amount    int
}

type knot struct {
	x int
	y int
}

type rope struct {
	head knot
	tail knot
}

type matrix [6000][6000]bool

func main() {
	input := input("input.txt")
	moves := parse(input)
	w, h := calcLimits(moves)
	println(w, h)

	m := matrix{}
	m.simulate(moves)
	count := m.countMoves()

	// m.print()
	println(count)
}

func (mat *matrix) simulate(moves []move) {
	halfway := len(mat) / 2
	r := rope{
		head: knot{halfway, halfway},
		tail: knot{halfway, halfway},
	}

	// Mark initial tail position in matrix.
	mat[r.tail.y][r.tail.x] = true

	for _, m := range moves {
		r.move(mat, m)
	}
}

func (mat *matrix) countMoves() int {
	count := 0

	for _, i := range mat {
		for _, j := range i {
			if j {
				count++
			}
		}
	}

	return count
}

func (mat *matrix) print() {
	for _, i := range mat {
		for _, j := range i {
			if j {
				print("#")
			} else {
				print("-")
			}
		}
		println()
	}
}

func (r *rope) move(mat *matrix, m move) {
	for i := 0; i < m.amount; i++ {

		// Move head.
		if m.direction == "R" {
			r.head.x++
		} else if m.direction == "L" {
			r.head.x--
		} else if m.direction == "U" {
			r.head.y--
		} else if m.direction == "D" {
			r.head.y++
		}

		// Get direction that tail needs to move in.
		d, g := r.direction()
		// Move tail.
		if !g {
			// Tail not required to move.
			continue
		}
		if d == 0 {
			r.tail.y--
		} else if d == 45 {
			r.tail.x++
			r.tail.y--
		} else if d == 90 {
			r.tail.x++
		} else if d == 135 {
			r.tail.x++
			r.tail.y++
		} else if d == 180 {
			r.tail.y++
		} else if d == 225 {
			r.tail.x--
			r.tail.y++
		} else if d == 270 {
			r.tail.x--
		} else if d == 315 {
			r.tail.x--
			r.tail.y--
		}

		// Mark tail position in matrix.
		mat[r.tail.y][r.tail.x] = true
	}
}

func (r *rope) direction() (int, bool) {
	direction := 0
	distanceGreaterThan1 := false

	if r.head.x-r.tail.x > 1 || r.head.x-r.tail.x < -1 || r.head.y-r.tail.y > 1 || r.head.y-r.tail.y < -1 {
		distanceGreaterThan1 = true
	}

	if r.head.x == r.tail.x && r.head.y == r.tail.y {
		// Head and tail overlap.
		return -1, false
	}

	if r.head.x == r.tail.x {
		if r.head.y < r.tail.y {
			direction = 0
		} else {
			direction = 180
		}
	}

	if r.head.y == r.tail.y {
		if r.head.x < r.tail.x {
			direction = 270
		} else {
			direction = 90
		}
	}

	if r.head.x > r.tail.x && r.head.y < r.tail.y {
		direction = 45
	} else if r.head.x < r.tail.x && r.head.y < r.tail.y {
		direction = 315
	} else if r.head.x > r.tail.x && r.head.y > r.tail.y {
		direction = 135
	} else if r.head.x < r.tail.x && r.head.y > r.tail.y {
		direction = 225
	}

	return direction, distanceGreaterThan1
}

// calcLimits returns the maximum potential size of the matrix.
func calcLimits(moves []move) (int, int) {
	l, r, u, d := 0, 0, 0, 0

	for _, m := range moves {
		if m.direction == "R" {
			r += m.amount
		} else if m.direction == "L" {
			l += m.amount
		} else if m.direction == "U" {
			u += m.amount
		} else if m.direction == "D" {
			d += m.amount
		}
	}

	maxwidth := l + r
	maxheight := u + d
	return maxwidth, maxheight
}

func parse(input []string) []move {
	moves := []move{}

	for _, l := range input {
		s := strings.Split(l, " ")
		a, _ := strconv.Atoi(s[1])
		m := move{
			direction: s[0],
			amount:    a,
		}

		moves = append(moves, m)
	}

	return moves
}

func input(name string) []string {
	bytes, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bytes), "\n")
}
