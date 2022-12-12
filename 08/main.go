package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("part one: %d\n", partOne())
	fmt.Printf("part two: %d\n", partTwo())
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

func partTwo() int {
	c := make(chan string)
	go readInput(c, "input.txt")

	g := grid(make([][]height, 0))
	for s := range c {
		row := parseRow(s)
		g = append(g, row)
	}

	var maxDistance int
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			vt := countVisibleTress(g, i, j)
			vd := calculateViewingDistance(vt)

			if vd > maxDistance {
				maxDistance = vd
			}
		}
	}

	return maxDistance
}

func countVisibleTress(g grid, i, j int) [4]int {
	return [4]int{
		countTopVisibleTrees(g, i, j),
		countLeftVisibleTrees(g, i, j),
		countRightVisibleTrees(g, i, j),
		countBottomVisibleTrees(g, i, j),
	}
}

func countTopVisibleTrees(g grid, initialRow, column int) int {
	h := g[initialRow][column]
	for row := initialRow - 1; row >= 0; row-- {
		if n := g[row][column]; n >= h {
			return initialRow - row
		}
	}

	return initialRow
}

func countLeftVisibleTrees(g grid, row, initialColumn int) int {
	h := g[row][initialColumn]
	for column := initialColumn - 1; column >= 0; column-- {
		if n := g[row][column]; n >= h {
			return initialColumn - column
		}
	}

	return initialColumn
}

func countRightVisibleTrees(g grid, row, initialColumn int) int {
	h := g[row][initialColumn]
	for column := initialColumn + 1; column < len(g[row]); column++ {
		if n := g[row][column]; n >= h {
			return column - initialColumn
		}
	}

	return len(g[row]) - 1 - initialColumn
}

func countBottomVisibleTrees(g grid, initialRow, column int) int {
	h := g[initialRow][column]
	for row := initialRow + 1; row < len(g); row++ {
		if n := g[row][column]; n >= h {
			return row - initialRow
		}
	}

	return len(g) - 1 - initialRow
}

func calculateViewingDistance(vt [4]int) int {
	d := 1
	for _, v := range vt {
		d *= v
	}
	return d
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
