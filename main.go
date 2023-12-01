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
	dirsPtr := flag.Bool("dirs", false, "Create directories and empty files for input")

	flag.Parse()

	fmt.Println("dirs:", *dirsPtr)

	if *dirsPtr == true {
		generateDirs()
	}
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
