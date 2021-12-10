package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Graph struct {
	root *Node
}

func (g *Graph) bfs() {
	fmt.Println("hello")
	queue := []*Node{}
	queue = append(queue, g.root)
	result := []int{}

	for len(queue) != 0 {
		result = append(result, queue[0].Value)

		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}

		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
	}

	for _, value := range result {
		fmt.Printf("%d ", value)
	}
}

func (graph *Graph) dfs() {

}

func main() {
	graph := &Graph{&Node{1, &Node{2, &Node{4, nil, nil}, nil}, &Node{3, nil, nil}}}
	graph.bfs()
}
