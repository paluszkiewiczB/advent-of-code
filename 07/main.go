package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("part one: %s\n", partOne())
	fmt.Printf("part two: %s\n", partTwo())
}

func partOne() string {
	c := make(chan string)
	go readInput(c, "sample-input.txt")
	for s := range c {
		println(s)
	}

	return "todo"
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
