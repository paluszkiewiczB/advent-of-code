package main

import (
	"bufio"
	"fmt"
	"os"
)

type shape = string

const (
	rock     shape = "rock"
	paper    shape = "paper"
	scissors shape = "scissors"
)

type move struct {
	opponent, my shape
}

var (
	shapePoints = map[shape]int{
		rock:     1,
		paper:    2,
		scissors: 3,
	}
)

type strategy struct {
	shape
	result
}

type result = string

const (
	lose result = "lose"
	draw result = "draw"
	win  result = "win"
)

func main() {
	m := readMoves()
	fmt.Printf("part one points: %d\n", partOne(m))

	s := readStrategy()
	fmt.Printf("part two points: %d\n", partTwo(s))
}

func partOne(m []move) int {
	var points int
	for _, move := range m {
		points += shapePoints[move.my]
		points += calculateMatchPoints(move)
	}

	return points
}

func partTwo(s []strategy) int {
	var points int
	for _, strat := range s {
		myShape := findMyShape(strat)
		points += shapePoints[myShape]
		points += calculateMatchPoints(move{
			opponent: strat.shape,
			my:       myShape,
		})
	}

	return points
}

var winningMoves = map[shape]shape{
	rock:     paper,
	paper:    scissors,
	scissors: rock,
}

var losingMoves = map[shape]shape{
	rock:     scissors,
	paper:    rock,
	scissors: paper,
}

func findMyShape(s strategy) shape {
	switch s.result {
	case draw:
		return s.shape
	case win:
		return winningMoves[s.shape]
	case lose:
		return losingMoves[s.shape]
	}

	panic(fmt.Errorf("unsupported result: %s", s.shape))
}

func calculateMatchPoints(m move) int {
	if m.my == m.opponent {
		return 3
	}

	if m.opponent == rock && m.my == paper {
		return 6
	}

	if m.opponent == scissors && m.my == rock {
		return 6
	}

	if m.opponent == paper && m.my == scissors {
		return 6
	}

	return 0
}

var opponentMoves = map[uint8]shape{
	'A': rock,
	'B': paper,
	'C': scissors,
}
var myMoves = map[uint8]shape{
	'X': rock,
	'Y': paper,
	'Z': scissors,
}

func readMoves() []move {
	scanner, closeFunc := readInput()
	defer closeFunc()

	moves := make([]move, 0)
	for scanner.Scan() {
		line := scanner.Text()
		opponent, ook := opponentMoves[line[0]]
		my, mok := myMoves[line[2]]
		if !ook || !mok {
			panic(fmt.Errorf("could not parse move: %s. found values: %s, %s", line, opponent, my))
		}
		move := move{opponent: opponent, my: my}
		moves = append(moves, move)
	}

	return moves
}

var results = map[uint8]result{
	'X': lose,
	'Y': draw,
	'Z': win,
}

func readStrategy() []strategy {
	scanner, closeFunc := readInput()
	defer closeFunc()

	strats := make([]strategy, 0)
	for scanner.Scan() {
		line := scanner.Text()
		m, mok := opponentMoves[line[0]]
		r, rok := results[line[2]]
		if !mok || !rok {
			panic(fmt.Errorf("could not read strategy: %s. found values: %s, %s", line, m, r))
		}
		strats = append(strats, strategy{
			shape:  m,
			result: r,
		})
	}

	return strats
}

func readInput() (s *bufio.Scanner, closeFunc func() error) {
	file, err := os.Open("input.txt")
	must(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner, file.Close
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
