package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("part one result: %d\n", partOne())
	fmt.Printf("part two result: %d\n", partTwo())
}

type item = rune

func partTwo() int {
	scanner, closeFunc := readInput()
	defer closeFunc()

	badges := make([]item, 0)
	group := [3]string{}
	var i, sum int
	for scanner.Scan() {
		line := scanner.Text()
		group[i] = line
		if i == 2 {
			badges = append(badges, findBadge(group))
			i = 0
		} else {
			i++
		}
	}

	for _, b := range badges {
		sum += calcPriority(b)
	}

	return sum
}

func findBadge(group [3]string) item {
	for _, r := range group[0] {
		if strings.ContainsRune(group[1], r) && strings.ContainsRune(group[2], r) {
			return r
		}
	}
	panic(fmt.Errorf("badge not found for group: %s", group))
}

func partOne() int {
	scanner, closeFunc := readInput()
	defer closeFunc()

	var sum int
	for scanner.Scan() {
		rucksack := scanner.Text()
		lCompartment, rCompartment := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]
		sharedItems := findShared(lCompartment, rCompartment)
		sum += calcPriorities(sharedItems)
	}

	return sum
}

func calcPriorities(items []item) int {
	var sum int
	for _, r := range items {
		sum += calcPriority(r)
	}

	return sum
}

func calcPriority(r item) int {
	if r >= 'a' && r <= 'z' {
		return 1 + int(r-'a')
	}

	if r >= 'A' && r <= 'Z' {
		return 27 + int(r-'A')
	}

	panic(fmt.Errorf("unsupported item: %s", string(r)))
}

func findShared(left, right string) []item {
	shared := make(map[item]struct{})
	uniqueLeft := make(map[item]struct{})

	for _, r := range left {
		uniqueLeft[r] = struct{}{}
	}

	for _, r := range right {
		if _, ok := uniqueLeft[r]; ok {
			shared[r] = struct{}{}
		}
	}

	out := make([]rune, len(shared))
	var i int
	for r := range shared {
		out[i] = r
		i++
	}
	return out
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
