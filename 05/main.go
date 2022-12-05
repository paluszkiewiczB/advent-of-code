package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("part one: %s\n", partOne())
}

type stack struct {
	s []rune
}

func newStack() *stack {
	s := make([]rune, 0)
	return &stack{s: s}
}

func (s *stack) String() string {
	b := strings.Builder{}
	b.WriteString("stack: {")
	for _, r := range s.s {
		b.WriteString(" ")
		b.WriteRune(r)
	}
	b.WriteString(" }")
	return b.String()
}

func (s *stack) pop() rune {
	r := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return r
}

func (s *stack) push(r rune) {
	s.s = append(s.s, r)
}

type move struct {
	count, from, to int
}

func partOne() string {
	c := make(chan string)
	go readInput(c)

	stacks := parseStacks(c)
	printStacks(stacks)

	m := make(chan move)
	go parseMoves(c, m)
	for mv := range m {
		for i := 0; i < mv.count; i++ {
			stacks[mv.to].push(stacks[mv.from].pop())
		}
	}

	sb := strings.Builder{}
	for _, s := range stacks {
		sb.WriteRune(s.pop())
	}
	return sb.String()
}

func printStacks(stacks []*stack) {
	for i, s := range stacks {
		fmt.Printf("stack #%d: %s\n", i, s.String())
	}
}

func parseMoves(lines chan string, m chan move) {
	for line := range lines {
		split := strings.Split(line, " ")
		m <- move{
			count: mustParse(split[1]),
			from:  mustParse(split[3]) - 1,
			to:    mustParse(split[5]) - 1,
		}
	}

	close(m)
}

func mustParse(s string) int {
	i, err := strconv.Atoi(s)
	must(err)
	return i
}

func parseStacks(c chan string) []*stack {
	stacksString := make([]string, 0)
	for in := range c {
		if in == "" {
			break
		}
		stacksString = append(stacksString, in)
	}

	whitespaces := regexp.MustCompile("\\s+")
	nums := whitespaces.Split(stacksString[len(stacksString)-1], -1)
	stacksCount, err := strconv.Atoi(nums[len(nums)-1])
	must(err)

	stacks := make([]*stack, stacksCount)
	for i := range stacks {
		stacks[i] = newStack()
	}

	for i := len(stacksString) - 2; i >= 0; i-- {
		stackLine := stacksString[i]
		for j, s := range stacks {

			elementIndex := 1 + j*4
			if elementIndex >= len(stackLine) {
				continue
			}

			if element := rune(stackLine[elementIndex]); element != ' ' {
				s.push(element)
			}
		}
	}

	return stacks
}

func readInput(c chan string) {
	file, err := os.Open("input.txt")
	must(err)
	defer mustf(file.Close)
	defer close(c)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		c <- scanner.Text()
	}
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
