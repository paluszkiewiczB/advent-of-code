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

type sections struct {
	from, to int
}

type pair struct {
	a, b sections
}

func partTwo() int {
	scanner, closeFunc := readInput()
	defer closeFunc()

	var count int
	for scanner.Scan() {
		l := scanner.Text()
		sections := strings.Split(l, ",")
		pair := parsePair(sections)
		if partiallyOverlaps(pair) {
			count++
		}
	}

	return count
}

func partiallyOverlaps(p pair) bool {
	l, r := p.a, p.b
	if r.from < l.from || (l.from == r.from && r.to > l.to) {
		l, r = r, l
	}

	return l.to >= r.from
}

func partOne() int {
	scanner, closeFunc := readInput()
	defer closeFunc()

	var count int
	for scanner.Scan() {
		l := scanner.Text()
		sections := strings.Split(l, ",")
		pair := parsePair(sections)
		if fullyOverlaps(pair) {
			count++
		}
	}

	return count
}

func fullyOverlaps(p pair) bool {
	l, r := p.a, p.b
	if r.from < l.from || (l.from == r.from && r.to > l.to) {
		l, r = r, l
	}

	return l.to >= r.to
}

func parsePair(p []string) pair {
	return pair{
		a: parseSections(p[0]),
		b: parseSections(p[1]),
	}
}

func parseSections(s string) sections {
	limits := strings.Split(s, "-")
	from, err := strconv.Atoi(limits[0])
	must(err)
	to, err := strconv.Atoi(limits[1])
	must(err)
	return sections{
		from: from,
		to:   to,
	}
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
