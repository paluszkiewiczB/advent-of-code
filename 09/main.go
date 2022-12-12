package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("part one: %d\n", partOne())
	fmt.Printf("part two: %d\n", partTwo())
}

type state struct {
	head, tail *point
}

type rope struct {
	knots [10]*point
}

func newRope(initial point) *rope {
	knots := [10]*point{}
	for i := 0; i < 10; i++ {
		knots[i] = &point{x: initial.x, y: initial.y}
	}
	return &rope{knots: knots}
}

func (s *state) String() string {
	table := [5][6]rune{}
	table[s.tail.x][s.tail.y] = 'T'
	table[s.head.x][s.head.y] = 'H'
	sb := strings.Builder{}
	for _, row := range table {
		for _, c := range row {
			symbol := c
			if c == 0 {
				symbol = '.'
			}
			sb.WriteRune(symbol)
			sb.WriteRune(' ')
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

type point struct {
	x, y int
}

type move = rune

const (
	U move = 'U'
	L move = 'L'
	R move = 'R'
	D move = 'D'
)

func partTwo() int {
	c := make(chan string)
	go readInput(c, "input.txt")

	m := make(chan move)
	go readMoves(c, m)

	r := newRope(point{x: 4, y: 0})

	visited := make(map[point]struct{})
	visit := func() {
		visited[*r.knots[9]] = struct{}{}
	}
	//
	visit()
	//for mv := range m {
	//	moveHead(s, mv)
	//	updateTail(s)
	//	visit()
	//}

	return len(visited)
}

func partOne() int {
	c := make(chan string)
	go readInput(c, "input.txt")

	m := make(chan move)
	go readMoves(c, m)

	s := &state{
		head: &point{x: 4, y: 0},
		tail: &point{x: 4, y: 0},
	}

	visited := make(map[point]struct{})
	visit := func() {
		visited[*s.tail] = struct{}{}
	}

	visit()
	for mv := range m {
		moveHead(s.head, mv)
		updateTail(s.head, s.tail)
		visit()
	}

	return len(visited)
}

func moveHead(p *point, mv move) {
	switch mv {
	case U:
		p.x--
	case L:
		p.y--
	case R:
		p.y++
	case D:
		p.x++
	}
}

func updateTail(prev, next *point) {
	// head and tail are in the same row
	if prev.x == next.x {
		if prev.y-next.y == 2 {
			next.y++
			return
		}

		if next.y-prev.y == 2 {
			next.y--
			return
		}
		return
	}

	// head and tail are in the same column
	if prev.y == next.y {
		if prev.x-next.x == 2 {
			next.x++
			return
		}

		if next.x-prev.x == 2 {
			next.x--
			return
		}
		return
	}

	// head and tail are on the diagonal
	if xDist := prev.x - next.x; xDist == 1 || xDist == -1 {
		if yDist := prev.y - next.y; yDist == 1 || yDist == -1 {
			return
		}
	}

	// head and tail must be in different rows/columns but not touching diagonally
	if prev.x > next.x {
		next.x++
	} else {
		next.x--
	}
	if prev.y > next.y {
		next.y++
	} else {
		next.y--
	}
}

var moves = map[rune]move{
	'U': U,
	'L': L,
	'R': R,
	'D': D,
}

func readMoves(c chan string, m chan move) {
	for line := range c {
		split := strings.Split(line, " ")
		direction := moves[rune(split[0][0])]
		times, err := strconv.Atoi(split[1])
		must(err)
		for i := 0; i < times; i++ {
			m <- direction
		}
	}

	close(m)
}

func readInput(c chan string, fileName string) {
	file, err := os.Open(fileName)
	must(err)
	defer mustf(file.Close)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		c <- scanner.Text()
	}
	close(c)
}

func mustf(f func() error) {
	err := f()
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
