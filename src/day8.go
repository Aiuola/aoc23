package main

import (
	"fmt"
	"os"
	"strings"
)

func day8PartOne(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	lines := strings.Split(string(dat), "\n")
	instructions := parseInstructions(lines[0])
	nodes := parseNodes(lines[2:])

	for _, node := range nodes {
		fmt.Printf("%s\n", node.val)
	}

	return len(instructions)
}

func parseInstructions(line string) []bool {
	return make([]bool, 0)
}

func parseNodes(lines []string) []*UnlinkedNode {
	return make([]*UnlinkedNode, 0)
}

func day8PartTwo(path string) int {
	return len(path)
}
