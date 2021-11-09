package main

type Stack struct {
	nodes []*Node
}

type Node struct {
	value string
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(n *Node) {
	s.nodes = append(s.nodes, n)
}

func (s *Stack) pop() {
	s.nodes = s.nodes[:len(s.nodes)-1]
}
