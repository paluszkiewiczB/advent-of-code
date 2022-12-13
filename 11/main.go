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

type monkey struct {
	number                int
	startingItems         []int
	operation             func(int) int
	test                  func(int) bool
	trueThrow, falseThrow int
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

	iC := make(chan int)
	inspectionsC := make(chan map[int]int)
	go func() {
		inspectionCount := make(map[int]int)
		for monkeyIndex := range iC {
			current := inspectionCount[monkeyIndex]
			inspectionCount[monkeyIndex] = current + 1
		}
		inspectionsC <- inspectionCount
	}()

	for i := 0; i < 20; i++ {
		for _, mo := range monkeys {
			for _, item := range mo.startingItems {
				worryLvl := mo.operation(item)
				worryLvl /= 3
				if mo.test(worryLvl) {
					monkeys[mo.trueThrow].startingItems = append(monkeys[mo.trueThrow].startingItems, worryLvl)
				} else {
					monkeys[mo.falseThrow].startingItems = append(monkeys[mo.falseThrow].startingItems, worryLvl)
				}
				iC <- mo.number
			}
			mo.startingItems = make([]int, 0)
		}

		//fmt.Printf("After round: %d\n", i)
		//for _, mo := range monkeys {
		//	fmt.Println(mo)
		//}
	}

	close(iC)

	max, preMax := 0, 0
	inspections := <-inspectionsC
	for _, count := range inspections {
		//fmt.Printf("Monkey: %d inspected items %d times\n", mo, count)
		if count > max {
			max, preMax = count, max
			continue
		}

		if count > preMax {
			preMax = count
		}
	}

	//fmt.Printf("max: %d, premax: %d\n", max, preMax)
	return max * preMax
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
	test := func(a int) bool {
		return a%divisibleBy == 0
	}

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
		test:          test,
		trueThrow:     trueMonkeyNo,
		falseThrow:    falseMonkeyNo,
	}
}

func mustParseInt(s string) int {
	v, err := strconv.Atoi(s)
	must(err)
	return v
}

func partTwo() string {
	c := make(chan string)
	go readInput(c, "sample-input.txt")
	for s := range c {
		println(s)
	}

	return "todo"
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
