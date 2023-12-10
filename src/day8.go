package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func day8PartOne(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	lines := strings.Split(string(dat), "\n")
	// Right = true / left = false
	directions := parseDirections(lines[0])
	nodes := parseNodes(lines[2:])
	sort.Sort(nodes)

	i := 0
	cont := 0
	for {
		i = executeDirections(directions, nodes, i)
		cont++
		fmt.Printf("After the %d execution %s\n", cont, nodes[i])
		if nodes[i].val == "ZZZ" {
			break
		}
	}

	return cont * len(directions)
}

func parseDirections(line string) Directions {
	directions := make([]Direction, len(line))
	for i, char := range line {
		if char == 'L' {
			continue
		}
		directions[i] = true
	}
	return directions
}

func parseNodes(lines []string) UnlinkedNodes {
	nodes := make([]*UnlinkedNode, len(lines))
	for i, line := range lines {
		nodes[i] = parseNode(line)
	}
	return nodes
}

func parseNode(line string) *UnlinkedNode {
	var name, right, left string

	parts := strings.Split(line, "=")

	currWord := make([]int32, 0)
	for _, char := range parts[0] {
		if char == ' ' {
			name = string(currWord)
			currWord = make([]int32, 0)
			break
		}
		currWord = append(currWord, char)
	}

	links := parts[1][2 : len(parts[1])-1]
	for _, char := range links {
		if char == ' ' {
			continue
		}
		if char == ',' {
			left = string(currWord)
			currWord = make([]int32, 0)
			continue
		}
		currWord = append(currWord, char)
	}
	right = string(currWord)

	return NewUnlinkedNode(name, right, left)
}

func executeDirections(directions Directions, nodes UnlinkedNodes, startingIndex int) int {
	currIndex := startingIndex
	var newIndex int

	for _, direction := range directions {
		newIndex = nodes.Move(currIndex, direction)
		if newIndex == len(nodes) {
			panic(fmt.Sprintf(
				"Couldn't move in direction %s from node %s",
				direction,
				nodes[currIndex],
			))
		}
		currIndex = newIndex
		/*fmt.Printf(
			"Moved to node %s - at index %d\n",
			nodes[currIndex],
			currIndex,
		)*/
	}
	return currIndex
}

func day8PartTwo(path string) int {
	return len(path)
}
