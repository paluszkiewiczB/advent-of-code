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

type rope struct {
	knots []*point
}

func newRope(initial point, size int) *rope {
	knots := make([]*point, size)
	for i := 0; i < size; i++ {
		knots[i] = &point{x: initial.x, y: initial.y}
	}
	return &rope{knots: knots}
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

	r := newRope(point{x: 4, y: 0}, 10)

	return moveTheRope(r, m)
}

func partOne() int {
	c := make(chan string)
	go readInput(c, "input.txt")

	m := make(chan move)
	go readMoves(c, m)

	r := newRope(point{x: 4, y: 0}, 2)

	return moveTheRope(r, m)
}

func moveTheRope(r *rope, m chan move) int {
	visited := make(map[point]struct{})
	visit := func() {
		visited[*r.knots[len(r.knots)-1]] = struct{}{}
	}

	visit()
	for mv := range m {
		moveHead(r.knots[0], mv)
		for i := 1; i < len(r.knots); i++ {
			updateTail(r.knots[i-1], r.knots[i])
		}
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

func updateTail(prev, current *point) {
	// head and tail are in the same row
	if prev.x == current.x {
		if prev.y-current.y == 2 {
			current.y++
			return
		}

		if current.y-prev.y == 2 {
			current.y--
			return
		}
		return
	}

	// head and tail are in the same column
	if prev.y == current.y {
		if prev.x-current.x == 2 {
			current.x++
			return
		}

		if current.x-prev.x == 2 {
			current.x--
			return
		}
		return
	}

	// head and tail are on the diagonal
	if xDist := prev.x - current.x; xDist == 1 || xDist == -1 {
		if yDist := prev.y - current.y; yDist == 1 || yDist == -1 {
			return
		}
	}

	// head and tail must be in different rows/columns but not touching diagonally
	if prev.x > current.x {
		current.x++
	} else {
		current.x--
	}
	if prev.y > current.y {
		current.y++
	} else {
		current.y--
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
