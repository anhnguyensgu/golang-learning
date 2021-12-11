package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	q := []*Node{}

	if root != nil {
		q = append(q, root)
	}

	for len(q) != 0 {
		temp := []*Node{}

		for i := 0; i < len(q); i++ {
			node := q[i]
			if node.Left != nil {
				temp = append(temp, node.Left)
			}

			if node.Right != nil {
				temp = append(temp, node.Right)
			}

			if i+1 < len(q) {
				node.Next = q[i+1]
			}
		}

		q = temp
	}

	return root
}

func main() {
	root := &Node{1, &Node{2, nil, nil, nil}, &Node{3, nil, nil, nil}, nil}

	connect(root)
}
