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
	//fmt.Printf("part two: %s\n", partTwo())
}

func partOne() int {
	c := make(chan string)
	go readInput(c, "input.txt")

	xC := make(chan int)
	sum := make(chan int)
	go func() {
		s := 0
		for v := range xC {
			s += v
		}
		sum <- s
	}()

	interestingCycles := map[int]struct{}{20: {}, 60: {}, 100: {}, 140: {}, 180: {}, 220: {}}
	cycle, x := 1, 1
	nextCycle := func() {
		if _, ok := interestingCycles[cycle]; ok {
			fmt.Printf("interesting cycle number: %d, value: %d\n", cycle, x)
			xC <- x * cycle
		}
		cycle++
	}

	for s := range c {
		if s == "noop" {
			nextCycle()
			continue
		}

		add := strings.Split(s, " ")
		v, err := strconv.Atoi(add[1])
		must(err)
		nextCycle()
		nextCycle()
		x += v
	}

	close(xC)
	return <-sum
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
