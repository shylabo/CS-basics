package main

type Queue struct {
	nodes []*Node
}

type Node struct {
	value string
}

func NewQueue() *Queue {
	return &Queue{}
}

func (s *Queue) enqueue(n *Node) {
	s.nodes = append(s.nodes, n)
}

func (s *Queue) dequeue() {
	s.nodes = s.nodes[1:len(s.nodes)]
}
