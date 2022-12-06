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

type marker struct {
	m map[rune]int
	r map[int]rune
}

func newMarker(size int) *marker {
	return &marker{
		m: make(map[rune]int, size),
		r: make(map[int]rune, size),
	}
}

func (m *marker) add(i int, r rune) {
	m.m[r] = i
	m.r[i] = r
}

func (m *marker) indexOf(r rune) int {
	return m.m[r]
}

func (m *marker) atIndex(i int) rune {
	return m.r[i]
}

func (m *marker) len() int {
	return len(m.m)
}

func (m *marker) String() string {
	sb := strings.Builder{}
	sb.WriteString("{ ")
	for i, r := range m.r {
		sb.WriteString(fmt.Sprintf("%d: %s ", i, string(r)))
	}
	sb.WriteString("}")
	return sb.String()
}

func partOne() int {
	c := make(chan rune)
	go readInput(c, "input.txt")
	pos := 1
	m := newMarker(4)
	for r := range c {
		if existing, ok := m.m[r]; ok {
			nm := newMarker(4)
			for i := existing + 1; i < pos; i++ {
				nm.add(i, m.atIndex(i))
			}
			m = nm
		}
		m.add(pos, r)
		if m.len() == 4 {
			return m.indexOf(r)
		}
		pos++
	}

	return -1
}

func partTwo() int {
	c := make(chan rune)
	go readInput(c, "input.txt")
	pos := 1
	m := newMarker(14)
	for r := range c {
		if existing, ok := m.m[r]; ok {
			nm := newMarker(14)
			for i := existing + 1; i < pos; i++ {
				nm.add(i, m.atIndex(i))
			}
			m = nm
		}
		m.add(pos, r)
		if m.len() == 14 {
			//fmt.Printf("size reached at pos: %d with marker: %s\n", pos, m.String())
			return m.indexOf(r)
		}
		pos++
	}

	return -1
}

func readInput(c chan rune, fileName string) {
	file, err := os.Open(fileName)
	must(err)
	defer mustf(file.Close)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for _, r := range scanner.Text() {
		c <- r
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
