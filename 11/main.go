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

type monkey struct {
	number                int
	startingItems         []int
	operation             func(int) int
	trueThrow, falseThrow int
	divisibleBy           int
}

func (m *monkey) String() string {
	return fmt.Sprintf("Monkey: %d %v", m.number, m.startingItems)
}

func partOne() int {
	c := make(chan string)
	go readInput(c, "input.txt")
	m := make(chan *monkey)
	go parseMonkeys(c, m)

	monkeys := make([]*monkey, 0)
	for mo := range m {
		monkeys = append(monkeys, mo)
	}

	modFactor := 1
	for _, mo := range monkeys {
		modFactor *= mo.divisibleBy
	}
	fmt.Printf("mod factor: %d\n", modFactor)

	inspections := make(map[int]int)
	inspect := func(mo, items int) {
		current := inspections[mo]
		inspections[mo] = current + items
	}

	for i := 0; i < 20; i++ {
		for _, mo := range monkeys {
			for _, item := range mo.startingItems {
				worryLvl := mo.operation(item)
				worryLvl /= 3
				worryLvl %= modFactor
				if worryLvl%mo.divisibleBy == 0 {
					monkeys[mo.trueThrow].startingItems = append(monkeys[mo.trueThrow].startingItems, worryLvl)
				} else {
					monkeys[mo.falseThrow].startingItems = append(monkeys[mo.falseThrow].startingItems, worryLvl)
				}
			}
			inspect(mo.number, len(mo.startingItems))
			mo.startingItems = make([]int, 0)
		}
	}

	return calculateMonkeyBusiness(inspections)
}

func partTwo() int {
	c := make(chan string)
	go readInput(c, "input.txt")
	m := make(chan *monkey)
	go parseMonkeys(c, m)

	monkeys := make([]*monkey, 0)
	for mo := range m {
		monkeys = append(monkeys, mo)
	}

	modFactor := 1
	for _, mo := range monkeys {
		modFactor *= mo.divisibleBy
	}
	fmt.Printf("mod factor: %d\n", modFactor)

	inspections := make(map[int]int)
	inspect := func(mo, items int) {
		current := inspections[mo]
		inspections[mo] = current + items
	}

	for i := 0; i < 10000; i++ {
		for _, mo := range monkeys {
			for _, item := range mo.startingItems {
				worryLvl := mo.operation(item)
				worryLvl %= modFactor
				if worryLvl%mo.divisibleBy == 0 {
					monkeys[mo.trueThrow].startingItems = append(monkeys[mo.trueThrow].startingItems, worryLvl)
				} else {
					monkeys[mo.falseThrow].startingItems = append(monkeys[mo.falseThrow].startingItems, worryLvl)
				}
			}
			inspect(mo.number, len(mo.startingItems))
			mo.startingItems = make([]int, 0)
		}

	}

	return calculateMonkeyBusiness(inspections)
}

func parseMonkeys(c <-chan string, m chan<- *monkey) {
	i := 0
	buff := [6]string{}
	for s := range c {
		buff[i] = s
		i++
		if i == 6 {
			m <- parseMonkey(buff)
			_, ok := <-c
			if !ok {
				close(m)
				return
			}
			i = 0
			buff = [6]string{}
		}
	}
}

func parseMonkey(buff [6]string) *monkey {
	title := buff[0]
	numS := title[len("Monkey ") : len(title)-1]
	num := mustParseInt(numS)

	itemsLine := buff[1]
	itemsString := itemsLine[len("  Starting items: "):]
	items := strings.Split(itemsString, ", ")
	startingItems := make([]int, len(items))
	for i, item := range items {
		startingItems[i] = mustParseInt(item)
	}

	operationLine := buff[2]
	operationString := operationLine[len("  Operation: new = old "):]
	operation := func(a int) int {
		ops := strings.Split(operationString, " ")
		operator := ops[0]
		operand := a
		if ops[1] != "old" {
			operand = mustParseInt(ops[1])
		}

		switch operator {
		case "+":
			return a + operand
		case "*":
			return a * operand
		}

		panic("unsupported operator: " + operator)
	}

	testLine := buff[3]
	testString := testLine[len("  Test: divisible by "):]
	divisibleBy := mustParseInt(testString)

	trueLine := buff[4]
	trueString := trueLine[len("    If true: throw to monkey "):]
	trueMonkeyNo := mustParseInt(trueString)

	falseLine := buff[5]
	falseString := falseLine[len("    If false: throw to monkey "):]
	falseMonkeyNo := mustParseInt(falseString)

	return &monkey{
		number:        num,
		startingItems: startingItems,
		operation:     operation,
		divisibleBy:   divisibleBy,
		trueThrow:     trueMonkeyNo,
		falseThrow:    falseMonkeyNo,
	}
}

func mustParseInt(s string) int {
	v, err := strconv.Atoi(s)
	must(err)
	return v
}

func calculateMonkeyBusiness(inspections map[int]int) int {
	max, preMax := 0, 0
	for _, count := range inspections {
		if count > max {
			max, preMax = count, max
			continue
		}

		if count > preMax {
			preMax = count
		}
	}

	return max * preMax
}

func readInput(c chan<- string, fileName string) {
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
