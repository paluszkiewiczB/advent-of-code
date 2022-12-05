package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("required parameter: day number")
	}
	dayArg := os.Args[1]
	d, err := strconv.Atoi(dayArg)
	if err != nil {
		log.Fatalf("day must be a number, passed: %s", dayArg)
	}

	day := fmt.Sprintf("%02d", d)
	log.Printf("creating day: %s", day)
	if stat, _ := os.Stat(day); stat != nil {
		log.Fatalf("directory already exists for day: %s", day)
	}

	err = os.Mkdir(day, os.ModeDir|os.ModeSetuid|os.ModeSetgid|0770)
	must(err)

	log.Printf("changing working directory to: %s", day)
	err = os.Chdir(day)
	must(err)

	log.Printf("copying template of main.go")
	mainS, err := os.Open("../template/main.go")
	must(err)
	defer mainS.Close()
	mainD, err := os.Create("main.go")
	must(err)
	defer mainD.Close()
	_, err = io.Copy(mainD, mainS)
	must(err)

	log.Printf("creating input.txt, sample-input.txt and task.md")
	_, err = os.Create("input.txt")
	must(err)
	_, err = os.Create("sample-input.txt")
	must(err)
	_, err = os.Create("task.md")
	must(err)

	log.Printf("initializing go module")
	err = exec.Command("go", "mod", "init", fmt.Sprintf("interactor.dev/advent-of-code/%s", day)).Run()
	must(err)
	err = exec.Command("go", "mod", "tidy").Run()
	must(err)

	log.Printf("committing directory: %s", day)
	err = exec.Command("git", "add", ".").Run()
	must(err)
	err = exec.Command("git", "commit", "-m", fmt.Sprintf("chore(%s): setting up day %s", day, day)).Run()
	must(err)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
