package main

import (
	"fmt"
	"sort"
)

type UnlinkedNode struct {
	val, right, left string
}

func (n UnlinkedNode) String() string {
	return fmt.Sprintf("[%s] right %s left %s", n.val, n.right, n.left)
}

func NewUnlinkedNode(val string, right string, left string) *UnlinkedNode {
	return &UnlinkedNode{val: val, right: right, left: left}
}

func (n UnlinkedNode) GetNeighbour(direction Direction) string {
	if direction {
		return n.right
	} else {
		return n.left
	}
}

type UnlinkedNodes []*UnlinkedNode

func (un UnlinkedNodes) Len() int {
	return len(un)
}

func (un UnlinkedNodes) Less(i, j int) bool {
	return un[i].val < un[j].val
}

func (un UnlinkedNodes) Swap(i, j int) {
	un[i], un[j] = un[j], un[i]
}

func (un UnlinkedNodes) String() string {
	ret := "Nodes:"
	for _, node := range un {
		ret = fmt.Sprintf("%s\n%s", ret, node)
	}
	return ret
}

func (un UnlinkedNodes) Move(i int, direction Direction) int {
	return un.Search(un[i].GetNeighbour(direction))
}

func (un UnlinkedNodes) Search(target string) int {
	return sort.Search(len(un), func(i int) bool {
		return un[i].val >= target
	})
}

type Direction bool

func (i Direction) String() string {
	if i {
		return "Right"
	}
	return "Left"
}

type Directions []Direction

func (i Directions) String() string {
	ret := "Directions:"
	for _, direction := range i {
		ret = fmt.Sprintf("%s\n%s", ret, direction)
	}
	return ret
}
