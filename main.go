package main

import (
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	allPtr := flag.Bool("all", false, "Run all days")
	dirsPtr := flag.Bool("dirs", false, "Create directories and empty files for input")
	dayPtr := flag.Int("day", 0, "Day to execute")
	flag.Parse()

	if *dirsPtr {
		generateDirs()
	}

	if *dayPtr == 0 {
		fmt.Println("No day specified please provide the day you want to execute")
		return
	}

	fmt.Printf("Chosen day %d\nRun all: %t\n", *dayPtr, *allPtr)
}

func day1() {
	dat, err := os.ReadFile("inputs/days/1/example.txt")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open("inputs/days/1/example.txt")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n\n", n1, string(b1[:n1]))
}
