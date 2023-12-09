package main

type UnlinkedNode struct {
	val, right, left string
}

func NewUnlinkedNode(val string, right string, left string) *UnlinkedNode {
	return &UnlinkedNode{val: val, right: right, left: left}
}

type Node struct {
	val         string
	right, left *Node
}

func NewNode(val string) *Node {
	return &Node{val: val, right: nil, left: nil}
}

func (n Node) link(r *Node, l *Node) {
	n.right = r
	n.left = l
}
