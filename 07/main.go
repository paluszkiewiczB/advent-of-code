package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("part one: %d\n", partOne())
	//fmt.Printf("part two: %s\n", partTwo())
}

type dir struct {
	name string
	size int
}

func partOne() int {
	c := make(chan string)
	go readInput(c, "input.txt")

	root := parseCd(<-c)

	fsC := make(chan dir)
	fs := make(map[string]int)
	go func() {
		for d := range fsC {
			if ex, ok := fs[d.name]; ok {
				fs[d.name] = ex + d.size
			} else {
				fs[d.name] = d.size
			}
		}
	}()

	parseDir(c, fsC, root)
	close(fsC)

	var total int
	for _, size := range fs {
		if size <= 100000 {
			total += size
		}
	}

	return total
}

func parseDir(c chan string, fs chan dir, wd string) {
	relativize := childOf(wd)
	for line := range c {
		if strings.HasPrefix(line, "$ ls") {
			line = parseLs(c, fs, wd)
		}

		if strings.HasPrefix(line, "$ cd ") {
			to := parseCd(line)

			if to == ".." {
				return
			}

			parseDir(c, fs, relativize(to))
		}
	}
}

func parseLs(c chan string, fs chan dir, wd string) string {
	relativize := childOf(wd)
	var dirSize int
	for line := range c {
		if strings.HasPrefix(line, "dir ") {
			d := strings.Split(line, "dir ")[1]
			fs <- dir{name: relativize(d)}
			continue
		}

		if strings.HasPrefix(line, "$") {
			fs <- dir{name: wd, size: dirSize}
			for _, parent := range parentDirs(wd) {
				fs <- dir{name: parent, size: dirSize}
			}
			return line
		}

		regularFile := strings.Split(line, " ")
		size, err := strconv.Atoi(regularFile[0])
		must(err)
		dirSize += size
	}

	fs <- dir{name: wd, size: dirSize}
	for _, parent := range parentDirs(wd) {
		fs <- dir{name: parent, size: dirSize}
	}
	return ""
}

func parentDirs(dir string) []string {
	out := make([]string, 0)
	for {
		dir = filepath.Dir(dir)
		if dir == "/" {
			return out
		}
		out = append(out, dir)
	}
}

func childOf(parent string) func(c string) string {
	return func(c string) string {
		if parent == "/" {
			return parent + c
		}
		return parent + "/" + c
	}
}

func parseCd(cd string) string {
	return strings.Split(cd, "$ cd ")[1]
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
