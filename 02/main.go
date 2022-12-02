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

func main() {
	m := readMoves()
	fmt.Printf("part one points: %d", partOne(m))
}

func partOne(m []move) int {
	var points int
	for _, move := range m {
		points += shapePoints[move.my]
		points += calculateMatchPoints(move)
	}

	return points
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
	file, err := os.Open("input.txt")
	defer file.Close()
	must(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

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

func must(err error) {
	if err != nil {
		panic(err)
	}
}
