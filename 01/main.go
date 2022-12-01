package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type top struct {
	max [3]int
}

func (t *top) accept(candidate int) {
	for i := range t.max {
		if candidate > t.max[i] {
			for j := len(t.max) - 1; j > i; j-- {
				t.max[j] = t.max[j-1]
			}
			t.max[i] = candidate
			break
		}
	}
}

func test() {
	t := &top{}
	t.accept(7)
	t.accept(1)
	t.accept(0)
	t.accept(-2)
	t.accept(2)
	if !(t.max[0] == 7 && t.max[1] == 2 && t.max[2] == 1) {
		log.Fatal("top does not work properly, was: ", t.max)
	}
}

func main() {
	test()
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("01/input.txt")
	must(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var max, current int

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if current > max {
				max = current
			}
			current = 0
			continue
		}

		cal, err := strconv.Atoi(line)
		must(err)

		current += cal
	}

	log.Println(max)
}

func partTwo() {
	file, err := os.Open("01/input.txt")
	must(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var t, current = &top{}, 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			t.accept(current)
			current = 0
			continue
		}

		cal, err := strconv.Atoi(line)
		must(err)

		current += cal
	}

	var sum int
	for _, v := range t.max {
		sum += v
	}

	log.Println(sum)
}

func must(err error) {
	if err != nil {
		log.Fatal("must", err)
	}
}
