package main

import (
	"bufio"
	"os"
)

func main() {
	println("hello world")
}

func readInput() (s *bufio.Scanner, closeFunc func() error) {
	file, err := os.Open("sample-input.txt")
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
