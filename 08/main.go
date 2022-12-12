package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("part one: %d\n", partOne())
	fmt.Printf("part two: %s\n", partTwo())
}

type height uint8
type grid [][]height

func (g grid) String() string {
	sb := strings.Builder{}
	for _, row := range g {
		for _, h := range row {
			sb.WriteRune(rune(h + '0'))
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func partOne() int {
	c := make(chan string)
	go readInput(c, "input.txt")

	g := grid(make([][]height, 0))
	for s := range c {
		row := parseRow(s)
		g = append(g, row)
	}

	var trees int

	for i := 1; i < len(g)-1; i++ { // skips first and last row
		for j := 1; j < len(g[i])-1; j++ { // skips first and last column
			if isVisible(g, i, j) {
				trees++
			}
		}
	}

	trees += 2 * len(g)    // top and bottom row
	trees += 2 * len(g[0]) // first and last column
	trees -= 4             // corners were counted twice
	return trees
}

func isVisible(g grid, i, j int) bool {
	return isVisibleFromTop(g, i, j) ||
		isVisibleFromLeft(g, i, j) ||
		isVisibleFromRight(g, i, j) ||
		isVisibleFromBottom(g, i, j)
}

func isVisibleFromBottom(g grid, initialRow, column int) bool {
	h := g[initialRow][column]
	tallest := height(0)
	for row := initialRow + 1; row < len(g); row++ {
		n := g[row][column]
		if n > tallest {
			tallest = n
		}
	}

	return h > tallest
}

func isVisibleFromRight(g grid, row, initialColumn int) bool {
	h := g[row][initialColumn]
	tallest := height(0)
	for column := initialColumn + 1; column < len(g[row]); column++ {
		n := g[row][column]
		if n > tallest {
			tallest = n
		}
	}

	return h > tallest
}

func isVisibleFromLeft(g grid, row, initialColumn int) bool {
	h := g[row][initialColumn]
	tallest := height(0)
	for column := initialColumn - 1; column >= 0; column-- {
		n := g[row][column]
		if n > tallest {
			tallest = n
		}
	}

	return h > tallest
}

func isVisibleFromTop(g grid, initialRow, column int) bool {
	h := g[initialRow][column]
	tallest := height(0)
	for row := initialRow - 1; row >= 0; row-- {
		n := g[row][column]
		if n > tallest {
			tallest = n
		}
	}

	return h > tallest
}

func parseRow(row string) []height {
	out := make([]height, len(row))

	for i := 0; i < len(row); i++ {
		h := row[i] - '0'
		out[i] = height(h)
	}

	return out
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
