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
	fmt.Printf("part two: %s\n", partTwo())
}

type state struct {
	head, tail *point
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
		moveHead(s, mv)
		updateTail(s)
		visit()
	}

	return len(visited)
}

func moveHead(s *state, mv move) {
	switch mv {
	case U:
		s.head.x--
	case L:
		s.head.y--
	case R:
		s.head.y++
	case D:
		s.head.x++
	}
}

func updateTail(s *state) {
	// head and tail are in the same row
	if s.head.x == s.tail.x {
		if s.head.y-s.tail.y == 2 {
			s.tail.y++
			return
		}

		if s.tail.y-s.head.y == 2 {
			s.tail.y--
			return
		}
		return
	}

	// head and tail are in the same column
	if s.head.y == s.tail.y {
		if s.head.x-s.tail.x == 2 {
			s.tail.x++
			return
		}

		if s.tail.x-s.head.x == 2 {
			s.tail.x--
			return
		}
		return
	}

	// head and tail are on the diagonal
	if xDist := s.head.x - s.tail.x; xDist == 1 || xDist == -1 {
		if yDist := s.head.y - s.tail.y; yDist == 1 || yDist == -1 {
			return
		}
	}

	// head and tail must be in different rows/columns but not touching diagonally
	if s.head.x > s.tail.x {
		s.tail.x++
	} else {
		s.tail.x--
	}
	if s.head.y > s.tail.y {
		s.tail.y++
	} else {
		s.tail.y--
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

func partTwo() string {
	c := make(chan string)
	go readInput(c, "sample-input.txt")
	for s := range c {
		println(s)
	}

	return "todo"
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
